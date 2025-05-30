package modify

import (
	"fmt"
	"os"
	"regexp"
)

var (
	REPLACEMENTS = []struct {
		Before *regexp.Regexp
		After  string
		File   string
	}{
		{
			// TODO: Update this for latest
			Before: regexp.MustCompile(`check_call\(\['pulumi', 'plugin', 'install', 'resource', 'spacelift', '[^']+']`),
			After:  `check_call(['pulumi', 'plugin', 'install', '--server', 'https:\/\/downloads.spacelift.io\/pulumi-plugins', 'resource', 'spacelift', '0.0.0']`,
			File:   "sdk/python/setup.py",
		},
		{
			// (?m) enables multiline mode, not working without it
			Before: regexp.MustCompile(`(?m)^import pulumi$`),
			After:  `import pulumi as pulumilib`,
			File:   "sdk/python/pulumi_spacelift/stack.py",
		},
		{
			Before: regexp.MustCompile(`pulumi.CustomResource`),
			After:  `pulumilib.CustomResource`,
			File:   "sdk/python/pulumi_spacelift/stack.py",
		},
		{
			Before: regexp.MustCompile(`pulumi.ResourceOptions`),
			After:  `pulumilib.ResourceOptions`,
			File:   "sdk/python/pulumi_spacelift/stack.py",
		},
		{
			Before: regexp.MustCompile(`pulumi.Input`),
			After:  `pulumilib.Input`,
			File:   "sdk/python/pulumi_spacelift/stack.py",
		},
		{
			Before: regexp.MustCompile(`pulumi.Output`),
			After:  `pulumilib.Output`,
			File:   "sdk/python/pulumi_spacelift/stack.py",
		},
		{
			Before: regexp.MustCompile(`@pulumi.getter`),
			After:  `@pulumilib.getter`,
			File:   "sdk/python/pulumi_spacelift/stack.py",
		},
		{
			Before: regexp.MustCompile(`@pulumi.input_type`),
			After:  `@pulumilib.input_type`,
			File:   "sdk/python/pulumi_spacelift/stack.py",
		},
		{
			Before: regexp.MustCompile(`pulumi.get`),
			After:  `pulumilib.get`,
			File:   "sdk/python/pulumi_spacelift/stack.py",
		},
		{
			Before: regexp.MustCompile(`pulumi.set`),
			After:  `pulumilib.set`,
			File:   "sdk/python/pulumi_spacelift/stack.py",
		},
	}
)

// ModifyPythonFiles modifies the python sdk so that it works with the spacelift properties
func ModifyPythonFiles() error {
	for _, replacement := range REPLACEMENTS {
		err := modifyFile(replacement.File, replacement.Before, replacement.After)
		if err != nil {
			return err
		}
	}
	return nil
}

func modifyFile(file string, before *regexp.Regexp, after string) error {
	fmt.Printf("Replacing %s with %s in %s\n", before, after, file)
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	data = before.ReplaceAll(data, []byte(after))
	os.WriteFile(file, data, 0644)
	return nil
}
