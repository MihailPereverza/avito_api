package service

import (
	"avito_api/internal/db/interface"
	"avito_api/internal/service/interface"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type ReportService struct {
	accountDB   db_interface.AccountDB
	operationDB db_interface.OperationDB
}

func NewReportService(
	dbAccount db_interface.AccountDB,
	dbOperation db_interface.OperationDB,
) service_interface.ReportServiceInterface {
	return &ReportService{
		dbAccount,
		dbOperation,
	}
}

func (r ReportService) GetReport() (string, error) {
	operations, err := r.operationDB.GetApproved()
	if err != nil {
		return "", err
	}

	t := time.Now()
	fileName := fmt.Sprintf("general_%d_%02d_%02d_%02d_%02d_%02d.csv",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)
	file, err := os.Create(fmt.Sprintf("./reports/%s", fileName))
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("cant open general csv")
	}

	writer := csv.NewWriter(file)
	var reportData = [][]string{{"ID операции (ID заказа)", "Название услуги", "ID пользователя", "Выручка", "Дата регистрации операции"}}
	for _, operation := range operations {
		reportData = append(reportData, []string{
			fmt.Sprintf("%d", operation.ID),
			fmt.Sprintf("%s", operation.ServiceTitle),
			fmt.Sprintf("%d", operation.AccountID),
			fmt.Sprintf("%f", operation.TotalCost),
			fmt.Sprintf("%s", operation.CreateTime),
		})
	}
	err = writer.WriteAll(reportData)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
