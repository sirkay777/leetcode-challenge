package implement_queue_using_stacks

import (
	"encoding/json"
	"testing"
)

func TestMyQueue(t *testing.T) {
	data := []struct {
		commands string
		params   string
		output   string
	}{
		{
			`["MyQueue","push","push","peek","pop","empty"]`,
			`[[],[1],[2],[],[],[]]`,
			`[null,null,null,1,1,false]`,
		},
	}
	var result []interface{}
	for _, d := range data {
		var commands []string
		err := json.Unmarshal([]byte(d.commands), &commands)
		if err != nil {
			t.Fatalf("invalid commands input: %s\n", d.commands)
		}
		var params [][]int
		err = json.Unmarshal([]byte(d.params), &params)
		if err != nil {
			t.Fatalf("invalid params input: %s\n", d.params)
		}
		var myQueue MyQueue
		for i := 0; i < len(commands); i++ {
			command := commands[i]
			param := -1
			if len(params[i]) > 0 {
				param = params[i][0]
			}
			switch command {
			case "MyQueue":
				myQueue = Constructor()
				result = append(result, nil)
			case "push":
				myQueue.Push(param)
				result = append(result, nil)
			case "peek":
				result = append(result, myQueue.Peek())
			case "pop":
				result = append(result, myQueue.Pop())
			case "empty":
				result = append(result, myQueue.Empty())
			}
		}
		output, err := json.Marshal(result)
		if err != nil {
			t.Fatal(err)
		}
		if string(output) != d.output {
			t.Errorf("expected: %s, got: %s", d.output, string(output))
		}
	}

}
