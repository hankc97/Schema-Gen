package lint

import (
	"strings"
	"github.com/hankc97/cmd/gen_schema"
)

func LintWorkflow(sink *ProblemSink, target *gen_schema.WorkflowNode) error {
	workflow := target.Value
	if err := lintWorkflowName(sink, workflow.Name); err != nil {
		return err
	}
	
	for _, job := range workflow.Jobs {
		if err := lintWorkflowJobs(sink, job); err != nil {
			return err
		}
	}
	return nil
}

func lintWorkflowName(sink *ProblemSink, target gen_schema.NameNode) error {
	name := target.Value
	if strings.HasPrefix(name, "HelloWorld") {
		sink.RecordProblem(target.Raw, "you're not allowed to say Hello World")
	}
	return nil
}

func lintWorkflowJobs(sink *ProblemSink, target gen_schema.JobNode) error {
	value := target.Value
	if len(value.Steps) > 0 && value.Uses != "" {
		sink.RecordProblem(target.Raw, `can't use "steps" with "uses"`)
	}
	return nil
}

// if err := lintWorkflowOn(sink, workflow.On); err != nil {
	// 	return err
	// }