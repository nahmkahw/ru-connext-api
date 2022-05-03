package student_repositories

func (r *studentRepoDB) Authentication(studentCode string) (token *PrepareTokenRepo, err error) {

	tempToken := PrepareTokenRepo{}
	query := `SELECT STD_CODE, (1) AS STATUS  FROM RUSMART.STUDENT WHERE STD_CODE = :param1`

	err = r.oracle_db.Get(&tempToken, query, studentCode)
	if err != nil {
		return nil, err
	}

	token = &tempToken

	return token, nil
}