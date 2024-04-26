package main

func main() {

}

func check1(next_action, accept_status, is_deleted int) bool {
	//((is_deleted = 2 AND accept_status = 3) OR next_action != 1)
	return next_action != 1 || (is_deleted == 2 && accept_status == 3)
}

func check2(next_action, accept_status, is_deleted int) bool {
	//(next_action != 1 OR is_deleted = 2) AND (next_action != 1 OR accept_status = 3)
	return !(next_action == 1 && is_deleted == 2) && !(next_action == 1 && accept_status == 3)
}

func check3(next_action, accept_status, is_deleted int) bool {
	//(next_action != 1 OR is_deleted = 2) AND (next_action != 1 OR accept_status = 3)
	return (next_action != 1 || is_deleted == 2) && (next_action != 1 || accept_status != 3)
}
