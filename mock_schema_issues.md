1. oneOf struct code flow
* Setting a pointer ( * ) instead will give us a nil value which we can use 
* going through each struct property 
* StdErrors: 
    * more than one format type 
    * 0 format types 

2. if there are not any types, get child $ref to init type
* The $ref gets us the type 
"oneOf": [
    {
        "$ref": "#/definitions/event"
    },

"event": {
    "$comment": "https://help.github.com/en/github/automating-your-workflow-with-github-actions/events-that-trigger-workflows",
    "type": "string",
    "enum": [
        "check_run",
        "check_suite",