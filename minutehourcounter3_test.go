package main

import "testing"

func TestRealConveyorQueue_AddToBack_Normal(t *testing.T) {
	// Arrange
	countQueue := 3
	rcq := NewRealConveyorQueue(countQueue)
	countNum := 5
	// Act
	rcq.AddToBack(countNum)
	// Assert
	if rcq.queue[0] != 0 {
		t.Errorf("rcq.AddToBack has somthing wrong, expected 0, got %d", rcq.queue[0])
	}
	if rcq.queue[1] != 0 {
		t.Errorf("rcq.AddToBack has somthing wrong, expected 0, got %d", rcq.queue[1])
	}
	if rcq.queue[2] != countNum {
		t.Errorf("rcq.AddToBack has somthing wrong, expected %d, got %d", countNum, rcq.queue[2])
	}
}

func TestRealConveyorQueue_AddToBack_QueueIsEmpty(t *testing.T) {
	// Arrange
	rcq := NewRealConveyorQueue(0)
	countNum := 5
	// Act
	rcq.AddToBack(countNum)
	// Assert
	if rcq.queue[0] != countNum {
		t.Errorf("rcq.AddToBack has something wrong, expected %d, got %d", countNum, rcq.queue[0])
	}
}

func TestRealConveyorQueue_TotalSum(t *testing.T) {
	// Arrange
	rcq := initial3RealConveyorQueueWith1_2_3()
	// Act
	act := rcq.TotalSum()
	// Assert
	if act != 6 {
		t.Errorf("rcq.TotalSum() has something wrong, expected 6, got %d", act)
	}
}

func TestRealConveyorQueue_Shift_Normal(t *testing.T) {
	// Arrange
	rcq := initial3RealConveyorQueueWith1_2_3()
	// Act
	rcq.Shift(2)
	// Assert
	if len(rcq.queue) != 3 {
		t.Errorf("rcq.Shift() has something wrong, expected 3, got %d", len(rcq.queue))
	}
	if rcq.queue[0] != 3 {
		t.Errorf("rcq.Shift() has something wrong, expected 3, got %d", rcq.queue[0])
	}
	if rcq.queue[1] != 0 {
		t.Errorf("rcq.Shift() has something wrong, expected 0, got %d", rcq.queue[1])
	}
	if rcq.queue[2] != 0 {
		t.Errorf("rcq.Shift() has something wrong, expected 0, got %d", rcq.queue[2])
	}
}

func TestRealConveyorQueue_Shift_EdgeIn(t *testing.T) {
	// Arrange
	rcq := initial3RealConveyorQueueWith1_2_3()
	// Act
	rcq.Shift(3)
	// Assert
	if len(rcq.queue) != 3 {
		t.Errorf("rcq.Shift() has something wrong, expected 3, got %d", len(rcq.queue))
	}
	if rcq.queue[0] != 0 {
		t.Errorf("rcq.Shift() has something wrong, expected 0, got %d", rcq.queue[0])
	}
	if rcq.queue[1] != 0 {
		t.Errorf("rcq.Shift() has something wrong, expected 0, got %d", rcq.queue[1])
	}
	if rcq.queue[2] != 0 {
		t.Errorf("rcq.Shift() has something wrong, expected 0, got %d", rcq.queue[2])
	}
}

func TestRealConveyorQueue_Shift_EdgeOut(t *testing.T) {
	// Arrange
	rcq := initial3RealConveyorQueueWith1_2_3()
	// Act
	rcq.Shift(4)
	// Assert
	if len(rcq.queue) != 3 {
		t.Errorf("rcq.Shift() has something wrong, expected 3, got %d", len(rcq.queue))
	}
	if rcq.queue[0] != 0 {
		t.Errorf("rcq.Shift() has something wrong, expected 0, got %d", rcq.queue[0])
	}
	if rcq.queue[1] != 0 {
		t.Errorf("rcq.Shift() has something wrong, expected 0, got %d", rcq.queue[1])
	}
	if rcq.queue[2] != 0 {
		t.Errorf("rcq.Shift() has something wrong, expected 0, got %d", rcq.queue[2])
	}
}

func initial3RealConveyorQueueWith1_2_3() *RealConveyorQueue {
	rcq := NewRealConveyorQueue(3)
	var val1, val2, val3, sum int
	val1, val2, val3 = 1, 2, 3
	for i, val := range []int{val1, val2, val3} {
		sum += val
		rcq.queue[i] += val
	}
	return rcq
}
