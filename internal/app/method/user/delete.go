package user

import (
	"fmt"
)

func (ur UserDBRepository) Delete (id int) error{
	ur.db = OpenDB(&ur.db)
	defer ur.db.Close()
	if err := ur.db.Ping(); err != nil{
		return err
	}
	stmt, err := ur.db.Query("UPDATE users SET first_name = 'deleted', last_name = 'deleted', email = 'deleted', encrypted_password = 'deleted' WHERE user_id = ?" ,id)
	HandleError(err)
	defer stmt.Close()
	fmt.Println("Deleted a user.")
	return nil
}
