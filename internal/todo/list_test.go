package todo

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Add a new item to the list", func(t *testing.T) {
		list := List{}
		list.Add("Test")
		if len(list) != 1 {
			t.Errorf("Expected 1 item in list, got %d", len(list))
		}
	})
}

func TestComplete(t *testing.T) {
	t.Run("Complete an item in the list", func(t *testing.T) {
		list := List{}
		list.Add("Test")
		list.Complete(0)
		if !list[0].Done {
			t.Errorf("Expected item to be completed")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete an item from the list", func(t *testing.T) {
		list := List{}
		list.Add("Test")
		list.Delete(0)
		if len(list) != 0 {
			t.Errorf("Expected 0 items in list, got %d", len(list))
		}
	})
}

func TestSaveGet(t *testing.T) {
	t.Run("Save and get a list", func(t *testing.T) {
		list := List{}
		list.Add("Test")
		list.Save("test.json")
		list.Get("test.json")
		if len(list) != 1 {
			t.Errorf("Expected 1 item in list, got %d", len(list))
		}
	})
}