1. If the reference's deeply nested property key(s) do not exist it will not create a struct
for the top level definition key
* i.e, properties: defaults -> $ref /defintions/defaults ->  
"defaults": {
            "type": "object",
            "properties": {
                "run": {
                    "type": "object",
                    "properties": {
                        "shell": {
                            "$ref": "#/definitions/shell"
                        },
                        "working-directory": {
                            "$ref": "#/definitions/working-directory"
                        }
                    },
                    "minProperties": 1,
                    "additionalProperties": false
                }
            },
            "minProperties": 1,
            "additionalProperties": false
        },
* If shell, working-directory definitions do not exist

2. anyOf, oneOf, allOf (create methods to handle these keywords)
* refers to type of input
* oneOf: normalJobs -> runs-on 
https://help.github.com/en/github/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#jobsjob_idruns-on

https://help.github.com/en/actions/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#self-hosted-runners

Either, array or string type 

3. if type == "object" must contain []*Schema, List of *Schema Structs 
* If it doesnt the output struct will be empty 

4. enum
* Create a field for this type in jsonschema.go
* confused on why "shell" has anyOf -> enums
https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#using-a-specific-shell
* possible change shell anyOf to oneOf? 

5. Pattern Properties function need to be made inside generator.go
--> properties -> jobs: has nested Pattern Properties 

6. parse workflow reference -> string splitting -> lacking tests
