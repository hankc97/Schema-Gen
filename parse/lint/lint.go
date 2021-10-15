package lint

import (
	"strings"
	"github.com/hankc97/fixtures/gen_schema"
)

func LintWorkflow(sink *ProblemSink, target *gen_schema.WorkflowRoot) error {
	workflow := target
	if err := lintWorkflowName(sink, workflow.Name); err != nil {
		return err
	}

	if err := lintWorkflowJobs(sink, workflow.Jobs.Value.Steps); err != nil {
		return err
	}

	// for _, step := range workflow.Jobs.Value.Steps.Value {
	// 	if err := lintWorkflowJobs(sink, step); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func lintWorkflowName(sink *ProblemSink, target gen_schema.NameRaw) error {
	name := target.Value
	if strings.HasPrefix(name, "HelloWorld") {
		sink.RecordProblem(target.Raw, "you're not allowed to say Hello World")
	}
	return nil
}

func lintWorkflowJobs(sink *ProblemSink, target gen_schema.StepsRaw) error {
	if len(target.Value) >= 4 {
		sink.RecordProblem(target.Raw, "can't use more than 3 steps")
	}
	return nil
}

// if len(value.Steps) > 0 && value.Uses != "" {
	// 	sink.RecordProblem(target.Raw, `can't use "steps" with "uses"`)
	// }

// if err := lintWorkflowOn(sink, workflow.On); err != nil {
	// 	return err
	// }