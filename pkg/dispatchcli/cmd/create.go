///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/vmware/dispatch/pkg/api/v1"
	"github.com/vmware/dispatch/pkg/dispatchcli/i18n"
	"github.com/vmware/dispatch/pkg/utils"
)

var (
	createLong = i18n.T(`Create a resource. See subcommands for resources that can be created.`)

	// TODO: Add examples
	createExample = i18n.T(``)
	file          = i18n.T(``)
	workDir       = i18n.T(``)
	baseURL       = i18n.T(``)
	isURL         = false
)

// ModelAction is the function type for CLI actions
type ModelAction func(interface{}) error

type importFunction struct {
	v1.Function
}

func resolveFileReference(ref string) (string, error) {
	if strings.HasPrefix(ref, "@") {
		filePath := path.Join(workDir, (ref)[1:])
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			return "", errors.Wrapf(err, "Error when reading content of %s", filePath)
		}
		encoded := string(fileContent)
		return encoded, nil
	}
	return ref, nil
}

func importFile(out io.Writer, errOut io.Writer, cmd *cobra.Command, args []string, actionMap map[string]ModelAction, actionName string) error {
	fullPath := path.Join(workDir, file)
	b, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return errors.Wrapf(err, "Error reading file %s", fullPath)
	}
	return importBytes(out, b, actionMap, actionName)
}

func importFileWithURL(out io.Writer, errOut io.Writer, cmd *cobra.Command, args []string, actionMap map[string]ModelAction, actionName string) error {
	resp, err := http.Get(file)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return importBytes(out, contents, actionMap, actionName)
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func importBytes(out io.Writer, b []byte, actionMap map[string]ModelAction, actionName string) error {

	var err error

	// Manually split up the yaml doc.  This is NOT a streaming parser.
	docs := bytes.Split(b, []byte("---"))

	type kind struct {
		Kind string `json:"kind"`
	}

	type output struct {
		APIs             []*v1.API             `json:"api"`
		BaseImages       []*v1.BaseImage       `json:"baseImages"`
		Images           []*v1.Image           `json:"images"`
		DriverTypes      []*v1.EventDriverType `json:"driverTypes"`
		Drivers          []*v1.EventDriver     `json:"drivers"`
		Subscriptions    []*v1.Subscription    `json:"subscriptions"`
		Functions        []*v1.Function        `json:"functions"`
		Secrets          []*v1.Secret          `json:"secrets"`
		Policies         []*v1.Policy          `json:"policies"`
		ServiceInstances []*v1.ServiceInstance `json:"serviceInstances"`
		ServiceAccounts  []*v1.ServiceAccount  `json:"serviceaccounts"`
		Organizations    []*v1.Organization    `json:"organizations"`
	}

	o := output{}

	for _, doc := range docs {
		k := &kind{}
		err = yaml.Unmarshal(doc, k)
		if err != nil {
			return errors.Wrapf(err, "Error decoding document %s", doc)
		}
		switch docKind := k.Kind; docKind {
		case utils.APIKind:
			m := &v1.API{}
			err = yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding api document %s", doc)
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.APIs = append(o.APIs, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.BaseImageKind:
			m := &v1.BaseImage{}
			err = yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding base image document %s", doc)
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.BaseImages = append(o.BaseImages, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.ImageKind:
			m := &v1.Image{}
			err = yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding image document %s", doc)
			}
			if m.RuntimeDependencies != nil {
				manifest, err := resolveFileReference(m.RuntimeDependencies.Manifest)
				if err != nil {
					return err
				}
				m.RuntimeDependencies.Manifest = manifest
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Images = append(o.Images, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.FunctionKind:
			m := &v1.Function{}
			if err := yaml.Unmarshal(doc, m); err != nil {
				return errors.Wrapf(err, "Error decoding function document %s", doc)
			}

			if m.SourcePath != "" {
				sourcePath := filepath.Join(workDir, m.SourcePath)
				isDir, err := utils.IsDir(sourcePath)
				if err != nil {
					if isURL {
						url := baseURL + m.SourcePath
						err = os.MkdirAll(sourcePath[:strings.LastIndex(sourcePath, "/")], 0755)
						downloadFile(sourcePath, url)
						if err != nil {
							return err
						}
					} else {
						return err
					}
				}
				if isDir && m.Handler == "" {
					return fmt.Errorf("error creating function %s: handler is required, source path %s is a directory", *m.Name, sourcePath)
				}
				sourceTarGz, err := utils.TarGzBytes(sourcePath)
				if err != nil {
					return errors.Wrapf(err, "Error when reading content of %s", sourcePath)
				}
				m.Source = sourceTarGz
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Functions = append(o.Functions, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.DriverTypeKind:
			m := &v1.EventDriverType{}
			err = yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error when decoding driver type document of %s", doc)
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.DriverTypes = append(o.DriverTypes, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.DriverKind:
			m := &v1.EventDriver{}
			err = yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding driver document %s", doc)
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Drivers = append(o.Drivers, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.SubscriptionKind:
			m := &v1.Subscription{}
			err = yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding subscription document %s", doc)
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Subscriptions = append(o.Subscriptions, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.SecretKind:
			m := &v1.Secret{}
			err = yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding secret document %s", doc)
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Secrets = append(o.Secrets, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.PolicyKind:
			m := &v1.Policy{}
			err = yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding policy document %s", doc)
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Policies = append(o.Policies, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.ServiceInstanceKind:
			m := &v1.ServiceInstance{}
			err := yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding service instance document %s", doc)
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.ServiceInstances = append(o.ServiceInstances, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.ServiceAccountKind:
			m := &v1.ServiceAccount{}
			err = yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding service account document %s", doc)
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.ServiceAccounts = append(o.ServiceAccounts, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		case utils.OrganizationKind:
			m := &v1.Organization{}
			err = yaml.Unmarshal(doc, m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding organization document %s", doc)
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Organizations = append(o.Organizations, m)
			fmt.Fprintf(out, "%s %s: %s\n", actionName, docKind, *m.Name)
		default:
			continue
		}
	}
	if dispatchConfig.JSON {
		encoder := json.NewEncoder(out)
		encoder.SetIndent("", "    ")
		return encoder.Encode(o)
	}
	return nil
}

var createMap map[string]ModelAction

func initCreateMap() {
	fnClient := functionManagerClient()
	imgClient := imageManagerClient()
	eventClient := eventManagerClient()
	apiClient := apiManagerClient()
	secClient := secretStoreClient()
	svcClient := serviceManagerClient()
	iamClient := identityManagerClient()

	createMap = map[string]ModelAction{
		utils.ImageKind:           CallCreateImage(imgClient),
		utils.BaseImageKind:       CallCreateBaseImage(imgClient),
		utils.FunctionKind:        CallCreateFunction(fnClient),
		utils.SecretKind:          CallCreateSecret(secClient),
		utils.ServiceInstanceKind: CallCreateServiceInstance(svcClient),
		utils.PolicyKind:          CallCreatePolicy(iamClient),
		utils.ApplicationKind:     CallCreateApplication,
		utils.ServiceAccountKind:  CallCreateServiceAccount(iamClient),
		utils.DriverTypeKind:      CallCreateEventDriverType(eventClient),
		utils.DriverKind:          CallCreateEventDriver(eventClient),
		utils.SubscriptionKind:    CallCreateSubscription(eventClient),
		utils.APIKind:             CallCreateAPI(apiClient),
		utils.OrganizationKind:    callCreateOrganization(iamClient),
	}
}

// NewCmdCreate creates a command object for the "create" action.
// Currently, one must use subcommands for specific resources to be created.
// In future create should accept file or stdin with multiple resources specifications.
func NewCmdCreate(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   i18n.T("Create resources."),
		Long:    createLong,
		Example: createExample,
		Run: func(cmd *cobra.Command, args []string) {
			if file == "" {
				runHelp(cmd, args)
				return
			}

			initCreateMap()
			if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
				isURL = true
				baseURL = file[:strings.LastIndex(file, "/")+1]
				err := importFileWithURL(out, errOut, cmd, args, createMap, "Created")
				CheckErr(err)
			} else {
				err := importFile(out, errOut, cmd, args, createMap, "Created")
				CheckErr(err)
			}

		},
	}

	cmd.Flags().StringVarP(&cmdFlagApplication, "application", "a", "", "associate with an application")
	cmd.Flags().StringVarP(&file, "file", "f", "", "Path to YAML file or an URL")
	cmd.Flags().StringVarP(&workDir, "work-dir", "w", "", "Working directory relative paths are based on")

	cmd.AddCommand(NewCmdCreateBaseImage(out, errOut))
	cmd.AddCommand(NewCmdCreateImage(out, errOut))
	cmd.AddCommand(NewCmdCreateFunction(out, errOut))
	cmd.AddCommand(NewCmdCreateSecret(out, errOut))
	cmd.AddCommand(NewCmdCreateAPI(out, errOut))
	cmd.AddCommand(NewCmdCreateSubscription(out, errOut))
	cmd.AddCommand(NewCmdCreateEventDriver(out, errOut))
	cmd.AddCommand(NewCmdCreateEventDriverType(out, errOut))
	cmd.AddCommand(NewCmdCreateApplication(out, errOut))
	cmd.AddCommand(NewCmdCreateServiceInstance(out, errOut))
	cmd.AddCommand(NewCmdCreateSeedImages(out, errOut))
	return cmd
}
