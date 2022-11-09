package hanlderReport

import (
	"avito_api/internal/config"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"log"
	"net/http"
)

// @Summary GetReport
// @Tags report
// @Description reserve balance from account
// @ID report-get-report
// @Accept json
// @Produce json
// @Success 200 {string} reportURL "reportURL"
// @failure 400 {object} outputModel.Error
// @Router /api/report/ [post]
func (rh *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	reportFileName, err := rh.reportService.GetReport()
	if err != nil {
		fmt.Printf("ReportHandler.GetReport error: %s \n", err)
		http.Error(w, "cant create report", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	reportURL := fmt.Sprintf("http://localhost:%s%s/%s", config.GetAppConfig().Port, config.GetAppConfig().ReportURI, reportFileName)

	resp := gabs.New()
	_, err = resp.Set("ok", "status")
	_, err = resp.Set(reportURL, "report_url")
	if err != nil {
		log.Fatalf("Error happened in set json. Err: %s", err)
	}
	if _, err := w.Write(resp.Bytes()); err != nil {
		log.Fatalf("Error happened in ResponseWriter Write. Err: %s", err)
	}
	return
}
