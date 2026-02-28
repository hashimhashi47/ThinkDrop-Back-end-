package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/pkg/constants"
)

func (s *ProfileService) RaiseComplaintService(req domain.CreateReportRequest, UserID uint) error {
	var Complaint domain.ReportComplaints

	Complaint = domain.ReportComplaints{
		Type:        req.Type,
		Description: req.Description,
		Status:      constants.PENDING,
		UserID:      UserID,
	}

	if err := s.repo.Create(&Complaint); err != nil {
		return errors.Join(err, errors.New("failed to request the complaint"))
	}

	return nil
}
