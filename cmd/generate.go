package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/whatDeepak/shellsage/helper/ai"
	"github.com/whatDeepak/shellsage/utils"
)

var commandTemplate = `
Imagine you are a security-conscious shell or terminal expert with a lot of computer knowledge.

Write a single, safe shell command that achieves the desired outcome. The command should:

* Not modify or delete files or folders.
* Be appropriate for a general audience (avoid offensive or harmful commands).
* Not require additional explanation.

Here is the prompt:

> %s

If the prompt is not related to a safe shell command or is not related to shell or commands or terminal, return "SHELLSAGE_AI_ERROR".

If the prompt is not appropriate for a general audience, return "SHELLSAGE_AI_ERROR".

If the prompt is unclear or ambiguous, return "SHELLSAGE_AI_ERROR".

If the prompt requires additional explanation, return "SHELLSAGE_AI_ERROR".

If the prompt is not a shell command, return "SHELLSAGE_AI_ERROR".

If the prompt is a shell command but is not safe, return "SHELLSAGE_AI_ERROR".

**Examples:**

* Prompt: "List all files in the current directory."
* Response: ls

* Prompt: "Delete all files in the current directory." (Unsafe)
* Response: "SHELLSAGE_AI_ERROR"

* Prompt: "Show a funny cat video." (Not a shell command)
* Response: "SHELLSAGE_AI_ERROR"
`

var rootCmd = &cobra.Command{
	Use:     "shellsage",
	Short:   "SHELLSAGE - Your AI-powered Shell Guide",
	Example: `SHELLSAGE "List all files in the current directory."`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := args[0]
		promptLength := len(prompt)

		if promptLength > 120 {
			utils.LogError("Prompt is too long. Please keep it under 120 characters.")
			return
		}

		spin := utils.GetSpinner()
		spin.Suffix = " Generating command..."
		spin.Start()
		resp, err := ai.Generate(cmd.Context(), commandTemplate, prompt)
		if err != nil {
			spin.Stop()
			utils.LogError(err.Error())
			return
		}

		spin.Stop()
		utils.LogInfo(resp)

		utils.HandleUserOptions(resp)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
