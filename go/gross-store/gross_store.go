package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill create a new bill
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem add item to customer bill
func AddItem(bill, units map[string]int, item string, unit string) bool {
	u, ok := units[unit]
	if ok {
		bill[item] = u
		return true
	}
	return false
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item string, unit string) bool {
	uVal, ok := units[unit]
	if !ok {
		return false
	}
	bVal, ok := bill[item]
	if !ok {
		return false
	}
	newVal := bVal - uVal
	if newVal < 0 {
		return false
	} else if newVal == 0 {
		delete(bill, item)
	} else {
		bill[item] = newVal
	}
	return true
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	el, ok := bill[item]
	if ok {
		return el, true
	}
	return 0, false
}
