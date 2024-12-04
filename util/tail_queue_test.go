package util

import (
	"testing"
)

func TestNewTailQueue(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		queue := NewTailQueue(5)
		if queue == nil {
			t.Errorf("Expected non-nil queue, got nil")
		}
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Invalid Size", func(t *testing.T) {
		queue := NewTailQueue(-1)
		if queue != nil {
			t.Errorf("Expected nil queue for invalid size, got non-nil")
		}
	})
}

func TestClear(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Clear()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0 after clear, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Already Empty", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Clear()
		queue.Clear()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0 after clear, got %d", queue.Count())
		}
	})
}

func TestCount(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0 for empty queue, got %d", queue.Count())
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if !queue.IsEmpty() {
			t.Errorf("Expected queue to be empty")
		}
	})

	t.Run("Happy Path - Non-Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.IsEmpty() {
			t.Errorf("Expected queue not to be empty")
		}
	})
}

func TestIsFull(t *testing.T) {
	t.Run("Happy Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		if !queue.IsFull() {
			t.Errorf("Expected queue to be full")
		}
	})

	t.Run("Happy Path - Not Full Queue", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		if queue.IsFull() {
			t.Errorf("Expected queue not to be full")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Wrap Around", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})
}

func TestHead(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Head() != 1 {
			t.Errorf("Expected head to be 1, got %v", queue.Head())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Head() != 1 {
			t.Errorf("Expected head to be 1, got %v", queue.Head())
		}
	})

	t.Run("Negative Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if queue.Head() != nil {
			t.Errorf("Expected head to be nil for empty queue")
		}
	})
}

func TestTail(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Tail() != 1 {
			t.Errorf("Expected tail to be 1, got %v", queue.Tail())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Tail() != 2 {
			t.Errorf("Expected tail to be 2, got %v", queue.Tail())
		}
	})

	t.Run("Negative Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if queue.Tail() != nil {
			t.Errorf("Expected tail to be nil for empty queue")
		}
	})
}

func TestKick(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0 for empty queue, got %d", queue.Count())
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if queue.Size() != 0 {
			t.Errorf("Expected size to be 0, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Size() != 1 {
			t.Errorf("Expected size to be 1, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Size() != 2 {
			t.Errorf("Expected size to be 2, got %d", queue.Size())
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if !queue.IsEmpty() {
			t.Errorf("Expected queue to be empty")
		}
	})

	t.Run("Happy Path - Non-Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.IsEmpty() {
			t.Errorf("Expected queue not to be empty")
		}
	})
}

func TestIsFull(t *testing.T) {
	t.Run("Happy Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		if !queue.IsFull() {
			t.Errorf("Expected queue to be full")
		}
	})

	t.Run("Happy Path - Not Full Queue", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		if queue.IsFull() {
			t.Errorf("Expected queue not to be full")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Wrap Around", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})
}

func TestKick(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0 for empty queue, got %d", queue.Count())
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if queue.Size() != 0 {
			t.Errorf("Expected size to be 0, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Size() != 1 {
			t.Errorf("Expected size to be 1, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Size() != 2 {
			t.Errorf("Expected size to be 2, got %d", queue.Size())
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if !queue.IsEmpty() {
			t.Errorf("Expected queue to be empty")
		}
	})

	t.Run("Happy Path - Non-Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.IsEmpty() {
			t.Errorf("Expected queue not to be empty")
		}
	})
}

func TestIsFull(t *testing.T) {
	t.Run("Happy Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		if !queue.IsFull() {
			t.Errorf("Expected queue to be full")
		}
	})

	t.Run("Happy Path - Not Full Queue", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		if queue.IsFull() {
			t.Errorf("Expected queue not to be full")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Wrap Around", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})
}

func TestKick(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0 for empty queue, got %d", queue.Count())
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if queue.Size() != 0 {
			t.Errorf("Expected size to be 0, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Size() != 1 {
			t.Errorf("Expected size to be 1, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Size() != 2 {
			t.Errorf("Expected size to be 2, got %d", queue.Size())
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if !queue.IsEmpty() {
			t.Errorf("Expected queue to be empty")
		}
	})

	t.Run("Happy Path - Non-Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.IsEmpty() {
			t.Errorf("Expected queue not to be empty")
		}
	})
}

func TestIsFull(t *testing.T) {
	t.Run("Happy Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		if !queue.IsFull() {
			t.Errorf("Expected queue to be full")
		}
	})

	t.Run("Happy Path - Not Full Queue", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		if queue.IsFull() {
			t.Errorf("Expected queue not to be full")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Wrap Around", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})
}

func TestKick(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0 for empty queue, got %d", queue.Count())
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if queue.Size() != 0 {
			t.Errorf("Expected size to be 0, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Size() != 1 {
			t.Errorf("Expected size to be 1, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Size() != 2 {
			t.Errorf("Expected size to be 2, got %d", queue.Size())
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if !queue.IsEmpty() {
			t.Errorf("Expected queue to be empty")
		}
	})

	t.Run("Happy Path - Non-Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.IsEmpty() {
			t.Errorf("Expected queue not to be empty")
		}
	})
}

func TestIsFull(t *testing.T) {
	t.Run("Happy Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		if !queue.IsFull() {
			t.Errorf("Expected queue to be full")
		}
	})

	t.Run("Happy Path - Not Full Queue", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		if queue.IsFull() {
			t.Errorf("Expected queue not to be full")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Wrap Around", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})
}

func TestKick(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0 for empty queue, got %d", queue.Count())
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if queue.Size() != 0 {
			t.Errorf("Expected size to be 0, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Size() != 1 {
			t.Errorf("Expected size to be 1, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Size() != 2 {
			t.Errorf("Expected size to be 2, got %d", queue.Size())
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if !queue.IsEmpty() {
			t.Errorf("Expected queue to be empty")
		}
	})

	t.Run("Happy Path - Non-Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.IsEmpty() {
			t.Errorf("Expected queue not to be empty")
		}
	})
}

func TestIsFull(t *testing.T) {
	t.Run("Happy Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		if !queue.IsFull() {
			t.Errorf("Expected queue to be full")
		}
	})

	t.Run("Happy Path - Not Full Queue", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		if queue.IsFull() {
			t.Errorf("Expected queue not to be full")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Wrap Around", func(t *testing.T) {
		queue := NewTailQueue(3)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Full Queue", func(t *testing.T) {
		queue := NewTailQueue(2)
		queue.Add(1)
		queue.Add(2)
		queue.Add(3)
		if queue.Count() != 2 {
			t.Errorf("Expected count to be 2, got %d", queue.Count())
		}
	})
}

func TestKick(t *testing.T) {
	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0, got %d", queue.Count())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		queue.Add(2)
		queue.Kick()
		if queue.Count() != 1 {
			t.Errorf("Expected count to be 1, got %d", queue.Count())
		}
	})

	t.Run("Negative Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Kick()
		if queue.Count() != 0 {
			t.Errorf("Expected count to be 0 for empty queue, got %d", queue.Count())
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("Happy Path - Empty Queue", func(t *testing.T) {
		queue := NewTailQueue(5)
		if queue.Size() != 0 {
			t.Errorf("Expected size to be 0, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Single Element", func(t *testing.T) {
		queue := NewTailQueue(5)
		queue.Add(1)
		if queue.Size() != 1 {
			t.Errorf("Expected size to be 1, got %d", queue.Size())
		}
	})

	t.Run("Happy Path - Multiple Elements", func(t *testing.T
