package api

import "testing"

func TestPatchedRunnerOptions(t *testing.T) {
	opts := DefaultOptions()

	err := opts.FromMap(map[string]any{
		"n_cpu_ffn":    float64(7),
		"tensor_split": "13,3",
		"kv_offload":   false,
	})
	if err != nil {
		t.Fatal(err)
	}

	if opts.NumCPUFFN == nil || *opts.NumCPUFFN != 7 {
		t.Fatalf("n_cpu_ffn not preserved: %v", opts.NumCPUFFN)
	}

	if opts.TensorSplit != "13,3" {
		t.Fatalf("tensor_split = %q", opts.TensorSplit)
	}

	if opts.KVOffload == nil || *opts.KVOffload {
		t.Fatalf("kv_offload false not preserved")
	}

	// Explicit zero must remain distinguishable from omitted.
	err = opts.FromMap(map[string]any{
		"n_cpu_ffn": float64(0),
	})
	if err != nil {
		t.Fatal(err)
	}

	if opts.NumCPUFFN == nil || *opts.NumCPUFFN != 0 {
		t.Fatal("explicit zero was not preserved")
	}

	// Integer values must not be accepted as booleans.
	if err := opts.FromMap(map[string]any{
		"kv_offload": float64(1),
	}); err == nil {
		t.Fatal("expected boolean validation error")
	}
}
