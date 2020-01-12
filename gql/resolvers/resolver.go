package resolvers

import (
	"context"

	"github.com/koba1108/gae-go-graphql-server/gql"
	"github.com/koba1108/gae-go-graphql-server/gql/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Test(ctx context.Context) (*string, error) {
	panic("not implemented")
}
func (r *mutationResolver) UploadRegistrationDocument(ctx context.Context, documentType *models.DocumentType, document *string) (*models.NacodoResponse, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateDocument(ctx context.Context, documentID *string, reviewStatus *models.DocumentReviewStatus, rejectReason *string) (*models.NacodoResponse, error) {
	panic("not implemented")
}
func (r *mutationResolver) ReplyToInquiry(ctx context.Context, inquiryID *string, text *string) (*models.NacodoResponse, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteInquiry(ctx context.Context, inquiryID *string) (*models.NacodoResponse, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePhoto(ctx context.Context, photoID *string, reviewStatus *models.PhotoReviewStatus, rejectReason *string) (*models.NacodoResponse, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, user *models.UserInput, photos *models.PhotosInput, details *models.UserDetailedProfileInput, tags []*models.TagInput) (*models.NacodoResponse, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Test(ctx context.Context) (*string, error) {
	panic("not implemented")
}
func (r *queryResolver) Documents(ctx context.Context, userID *string, typeArg *models.DocumentReviewStatus) ([]*models.Document, error) {
	panic("not implemented")
}
func (r *queryResolver) Favorites(ctx context.Context, id *string) (*models.UsersResponse, error) {
	panic("not implemented")
}
func (r *queryResolver) Inquiries(ctx context.Context, option *models.InquirySearchInput) ([]*models.Inquiry, error) {
	panic("not implemented")
}
func (r *queryResolver) Liked(ctx context.Context, id *string) ([]*models.Liked, error) {
	panic("not implemented")
}
func (r *queryResolver) Likes(ctx context.Context, id *string) ([]*models.Likes, error) {
	panic("not implemented")
}
func (r *queryResolver) Notifications(ctx context.Context, option *models.NotificationSearchInput) ([]*models.Notification, error) {
	panic("not implemented")
}
func (r *queryResolver) RegisterNotification(ctx context.Context, userIds []*string, notification *models.NotificationInput) (*models.NacodoResponse, error) {
	panic("not implemented")
}
func (r *queryResolver) UpdateNotification(ctx context.Context, userIds []*string, notification *models.NotificationInput) (*models.NacodoResponse, error) {
	panic("not implemented")
}
func (r *queryResolver) DeleteNotification(ctx context.Context, notificationID *string) (*models.NacodoResponse, error) {
	panic("not implemented")
}
func (r *queryResolver) Photos(ctx context.Context, userID *string, typeArg *models.PhotoReviewStatus) ([]*models.UserPhotos, error) {
	panic("not implemented")
}
func (r *queryResolver) Recommends(ctx context.Context, userID *string) (*models.UsersResponse, error) {
	panic("not implemented")
}
func (r *queryResolver) Reports(ctx context.Context, option *models.ReportSearchInput) ([]*models.Report, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id *string) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context, option *models.SearchOption, includes []*string, excludes []*string, sortKey *models.UserSortKey, numOfResults *int, offset *int, orderBy *models.OrderBy) (*models.UsersResponse, error) {
	panic("not implemented")
}
