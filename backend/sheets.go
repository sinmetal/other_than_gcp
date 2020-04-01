package backend

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/api/sheets/v4"
)

func SheetsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sheetID := r.FormValue("sheetID")

	s := NewSheetsService(ctx)
	status, err := s.ReadTest(ctx, sheetID, "A1")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Printf("failed write to response.err=%+v\n", err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(strconv.Itoa(status)))
	if err != nil {
		fmt.Printf("failed write to response.err=%+v\n", err)
	}
}

type SheetsService struct {
	svs *sheets.SpreadsheetsValuesService
}

func NewSheetsService(ctx context.Context) *SheetsService {
	ss, err := sheets.NewService(ctx)
	if err != nil {
		panic(err)
	}
	svs := sheets.NewSpreadsheetsValuesService(ss)
	return &SheetsService{
		svs: svs,
	}
}

func (s *SheetsService) ReadTest(ctx context.Context, sheetID string, sheetRange string) (int, error) {
	resp, err := s.svs.Get(sheetID, sheetRange).Do()
	if err != nil {
		return 0, err
	}

	return resp.HTTPStatusCode, nil
}
