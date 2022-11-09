package account

func (D *DBAccount) IsExists(accountID int) (bool, error) {
	exists := false
	query := `SELECT EXISTS(
    SELECT account_id FROM account
    WHERE account_id = $1
    )`
	err := D.db.QueryRow(
		query, accountID,
	).Scan(&exists)
	return exists, err
}
