package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	DomainAdmin "thinkdrop-backend/internal/modules/admin/domain"
)

// -> get all flaged posts
func (a *AdminService) GetAllFlagedPostService(limit, offset int) (interface{}, error) {
	var Post []domain.Post

	var PostResponse []DomainAdmin.AdminFlaggedPostResponse

	if err := a.repo.FindReportedPosts(&Post, limit, offset, "User", "Reports"); err != nil {
		return nil, errors.New("failed to find the posts")
	}

	for _, v := range Post {
		Reports := make([]DomainAdmin.AdminReportDTO, 0, len(v.Reports))
		for _, si := range v.Reports {
			Reports = append(Reports, DomainAdmin.AdminReportDTO{
				ID:          si.ID,
				ReportedBy:  si.UserID,
				Reason:      si.Reason,
				Description: si.Description,
				CreatedAt:   si.CreatedAt,
			})
		}
		PostResponse = append(PostResponse, DomainAdmin.AdminFlaggedPostResponse{
			ID:          v.ID,
			Content:     v.Content,
			CreatedAt:   v.CreatedAt,
			IsBlocked:   v.Blocked,
			ReportCount: v.ReportCount,
			User: DomainAdmin.AdminPostUserDTO{
				ID:            v.UserID,
				Name:          v.User.FullName,
				AnonymousName: v.User.AnonymousName,
				AvatarURL:     v.User.ImageURL,
			},
			Reports: Reports,
		})
	}
	return PostResponse, nil
}

func (a *AdminService) RemoveTheFlaggedPostService(PostID int) error {

	if err := a.repo.DeletePostWithRelations(uint(PostID)); err != nil {
		return errors.New("failed to delete the post")
	}

	return nil
}

func (a *AdminService) SafePostingService(PostID int) (interface{}, error) {
	var post domain.Post

	if err := a.repo.UpdateColumn(&post, "id = ?", PostID, "ReportCount", 0); err != nil {
		return nil, errors.New("failed to find the post")
	}

	return post, nil
}

func (a *AdminService) GetAllComplaintsService(limit, offset int) (interface{}, int, error) {
	var complaint []domain.ReportComplaints

	Total, _ := a.repo.Count(&domain.ReportComplaints{})
	if err := a.repo.FindAllWithOnePreload(&complaint, limit, offset, "User"); err != nil {
		return nil, 0, err
	}

	return complaint, int(Total), nil
}

func (a *AdminService) ConsiderTheIssueService(Postid int, Req DomainAdmin.UpdateComplaintStatusRequest) error {

	if err := a.repo.UpdateColumn(&domain.ReportComplaints{}, "id = ?", Postid,
		"Status", Req.Status); err != nil {
		return errors.New("failed to accept the post")
	}
	return nil
}
