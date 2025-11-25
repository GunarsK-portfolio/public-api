package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/GunarsK-portfolio/public-api/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// =============================================================================
// Test Constants
// =============================================================================

const (
	testProfileName    = "John Doe"
	testProfileTitle   = "Senior Software Engineer"
	testCertName       = "AWS Solutions Architect"
	testCertIssuer     = "Amazon Web Services"
	testSkillName      = "Go"
	testCompanyName    = "Acme Corp"
	testPosition       = "Senior Developer"
	testProjectName    = "Portfolio Website"
	testMiniatureName  = "Space Marine Captain"
	testMiniatureTheme = "Warhammer 40K"
)

// =============================================================================
// Mock Repository
// =============================================================================

type mockRepository struct {
	getProfileFunc              func(ctx context.Context) (*models.Profile, error)
	getAllWorkExperienceFunc    func(ctx context.Context) ([]models.WorkExperience, error)
	getAllCertificationsFunc    func(ctx context.Context) ([]models.Certification, error)
	getAllSkillsFunc            func(ctx context.Context) ([]models.Skill, error)
	getAllProjectsFunc          func(ctx context.Context) ([]models.PortfolioProject, error)
	getProjectByIDFunc          func(ctx context.Context, id int64) (*models.PortfolioProject, error)
	getAllMiniatureProjectsFunc func(ctx context.Context) ([]models.MiniatureProject, error)
	getMiniatureProjectByIDFunc func(ctx context.Context, id int64) (*models.MiniatureProject, error)
	getAllMiniatureThemesFunc   func(ctx context.Context) ([]models.MiniatureTheme, error)
}

func (m *mockRepository) GetProfile(ctx context.Context) (*models.Profile, error) {
	if m.getProfileFunc != nil {
		return m.getProfileFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetAllWorkExperience(ctx context.Context) ([]models.WorkExperience, error) {
	if m.getAllWorkExperienceFunc != nil {
		return m.getAllWorkExperienceFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetAllCertifications(ctx context.Context) ([]models.Certification, error) {
	if m.getAllCertificationsFunc != nil {
		return m.getAllCertificationsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetAllSkills(ctx context.Context) ([]models.Skill, error) {
	if m.getAllSkillsFunc != nil {
		return m.getAllSkillsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetAllProjects(ctx context.Context) ([]models.PortfolioProject, error) {
	if m.getAllProjectsFunc != nil {
		return m.getAllProjectsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetProjectByID(ctx context.Context, id int64) (*models.PortfolioProject, error) {
	if m.getProjectByIDFunc != nil {
		return m.getProjectByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetAllMiniatureProjects(ctx context.Context) ([]models.MiniatureProject, error) {
	if m.getAllMiniatureProjectsFunc != nil {
		return m.getAllMiniatureProjectsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetMiniatureProjectByID(ctx context.Context, id int64) (*models.MiniatureProject, error) {
	if m.getMiniatureProjectByIDFunc != nil {
		return m.getMiniatureProjectByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetAllMiniatureThemes(ctx context.Context) ([]models.MiniatureTheme, error) {
	if m.getAllMiniatureThemesFunc != nil {
		return m.getAllMiniatureThemesFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

// =============================================================================
// Test Helpers
// =============================================================================

func setupTestHandler(t *testing.T) (*Handler, *mockRepository) {
	t.Helper()

	gin.SetMode(gin.TestMode)
	mockRepo := &mockRepository{}
	handler := New(mockRepo)

	return handler, mockRepo
}

func setupTestRouter(t *testing.T) *gin.Engine {
	t.Helper()
	return gin.New()
}

func createTestProfile() models.Profile {
	return models.Profile{
		ID:        1,
		FullName:  testProfileName,
		Title:     testProfileTitle,
		Bio:       "Experienced software engineer",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func createTestCertification() models.Certification {
	return models.Certification{
		ID:        1,
		Name:      testCertName,
		Issuer:    testCertIssuer,
		IssueDate: "2024-01-15",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func createTestSkill() models.Skill {
	return models.Skill{
		ID:           1,
		Skill:        testSkillName,
		SkillTypeID:  1,
		IsVisible:    true,
		DisplayOrder: 1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func createTestWorkExperience() models.WorkExperience {
	return models.WorkExperience{
		ID:          1,
		Company:     testCompanyName,
		Position:    testPosition,
		Description: "Building awesome software",
		StartDate:   "2020-01-01",
		IsCurrent:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func createTestProject() models.PortfolioProject {
	return models.PortfolioProject{
		ID:          1,
		Title:       testProjectName,
		Description: "A portfolio website built with Vue and Go",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func createTestMiniatureProject() models.MiniatureProject {
	return models.MiniatureProject{
		ID:        1,
		Title:     testMiniatureName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func createTestMiniatureTheme() models.MiniatureTheme {
	return models.MiniatureTheme{
		ID:        1,
		Name:      testMiniatureTheme,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func performRequest(t *testing.T, router *gin.Engine, method, path string, body interface{}) *httptest.ResponseRecorder {
	t.Helper()

	var reqBody []byte
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}
		reqBody = b
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// =============================================================================
// Constructor Tests
// =============================================================================

func TestNewHandler(t *testing.T) {
	mockRepo := &mockRepository{}
	handler := New(mockRepo)

	if handler == nil {
		t.Error("New() should return non-nil handler")
	}
}

// =============================================================================
// Profile Handler Tests
// =============================================================================

func TestGetProfile_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/profile", handler.GetProfile)

	expectedProfile := createTestProfile()
	mockRepo.getProfileFunc = func(ctx context.Context) (*models.Profile, error) {
		return &expectedProfile, nil
	}

	w := performRequest(t, router, "GET", "/profile", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetProfile() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result models.Profile
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result.FullName != testProfileName {
		t.Errorf("GetProfile() name = %s, want %s", result.FullName, testProfileName)
	}
}

func TestGetProfile_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/profile", handler.GetProfile)

	mockRepo.getProfileFunc = func(ctx context.Context) (*models.Profile, error) {
		return nil, gorm.ErrRecordNotFound
	}

	w := performRequest(t, router, "GET", "/profile", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("GetProfile() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestGetProfile_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/profile", handler.GetProfile)

	mockRepo.getProfileFunc = func(ctx context.Context) (*models.Profile, error) {
		return nil, errors.New("database connection failed")
	}

	w := performRequest(t, router, "GET", "/profile", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetProfile() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

// =============================================================================
// Work Experience Handler Tests
// =============================================================================

func TestGetWorkExperience_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/experience", handler.GetWorkExperience)

	expectedExps := []models.WorkExperience{createTestWorkExperience()}
	mockRepo.getAllWorkExperienceFunc = func(ctx context.Context) ([]models.WorkExperience, error) {
		return expectedExps, nil
	}

	w := performRequest(t, router, "GET", "/experience", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetWorkExperience() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result []models.WorkExperience
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("GetWorkExperience() returned %d items, want 1", len(result))
	}
}

func TestGetWorkExperience_Empty(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/experience", handler.GetWorkExperience)

	mockRepo.getAllWorkExperienceFunc = func(ctx context.Context) ([]models.WorkExperience, error) {
		return []models.WorkExperience{}, nil
	}

	w := performRequest(t, router, "GET", "/experience", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetWorkExperience() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetWorkExperience_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/experience", handler.GetWorkExperience)

	mockRepo.getAllWorkExperienceFunc = func(ctx context.Context) ([]models.WorkExperience, error) {
		return nil, errors.New("database error")
	}

	w := performRequest(t, router, "GET", "/experience", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetWorkExperience() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

// =============================================================================
// Certifications Handler Tests
// =============================================================================

func TestGetCertifications_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/certifications", handler.GetCertifications)

	expectedCerts := []models.Certification{createTestCertification()}
	mockRepo.getAllCertificationsFunc = func(ctx context.Context) ([]models.Certification, error) {
		return expectedCerts, nil
	}

	w := performRequest(t, router, "GET", "/certifications", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetCertifications() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result []models.Certification
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("GetCertifications() returned %d items, want 1", len(result))
	}
}

func TestGetCertifications_Empty(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/certifications", handler.GetCertifications)

	mockRepo.getAllCertificationsFunc = func(ctx context.Context) ([]models.Certification, error) {
		return []models.Certification{}, nil
	}

	w := performRequest(t, router, "GET", "/certifications", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetCertifications() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetCertifications_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/certifications", handler.GetCertifications)

	mockRepo.getAllCertificationsFunc = func(ctx context.Context) ([]models.Certification, error) {
		return nil, errors.New("database error")
	}

	w := performRequest(t, router, "GET", "/certifications", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetCertifications() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

// =============================================================================
// Skills Handler Tests
// =============================================================================

func TestGetSkills_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/skills", handler.GetSkills)

	expectedSkills := []models.Skill{createTestSkill()}
	mockRepo.getAllSkillsFunc = func(ctx context.Context) ([]models.Skill, error) {
		return expectedSkills, nil
	}

	w := performRequest(t, router, "GET", "/skills", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetSkills() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result []models.Skill
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("GetSkills() returned %d items, want 1", len(result))
	}
}

func TestGetSkills_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/skills", handler.GetSkills)

	mockRepo.getAllSkillsFunc = func(ctx context.Context) ([]models.Skill, error) {
		return nil, errors.New("database error")
	}

	w := performRequest(t, router, "GET", "/skills", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetSkills() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

// =============================================================================
// Projects Handler Tests
// =============================================================================

func TestGetProjects_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/projects", handler.GetProjects)

	expectedProjects := []models.PortfolioProject{createTestProject()}
	mockRepo.getAllProjectsFunc = func(ctx context.Context) ([]models.PortfolioProject, error) {
		return expectedProjects, nil
	}

	w := performRequest(t, router, "GET", "/projects", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetProjects() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result []models.PortfolioProject
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("GetProjects() returned %d items, want 1", len(result))
	}
}

func TestGetProjects_Empty(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/projects", handler.GetProjects)

	mockRepo.getAllProjectsFunc = func(ctx context.Context) ([]models.PortfolioProject, error) {
		return []models.PortfolioProject{}, nil
	}

	w := performRequest(t, router, "GET", "/projects", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetProjects() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetProjects_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/projects", handler.GetProjects)

	mockRepo.getAllProjectsFunc = func(ctx context.Context) ([]models.PortfolioProject, error) {
		return nil, errors.New("database error")
	}

	w := performRequest(t, router, "GET", "/projects", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetProjects() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

func TestGetProjectByID_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/projects/:id", handler.GetProjectByID)

	expectedProject := createTestProject()
	mockRepo.getProjectByIDFunc = func(ctx context.Context, id int64) (*models.PortfolioProject, error) {
		if id != 1 {
			return nil, gorm.ErrRecordNotFound
		}
		return &expectedProject, nil
	}

	w := performRequest(t, router, "GET", "/projects/1", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetProjectByID() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result models.PortfolioProject
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result.Title != testProjectName {
		t.Errorf("GetProjectByID() title = %s, want %s", result.Title, testProjectName)
	}
}

func TestGetProjectByID_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/projects/:id", handler.GetProjectByID)

	mockRepo.getProjectByIDFunc = func(ctx context.Context, id int64) (*models.PortfolioProject, error) {
		return nil, gorm.ErrRecordNotFound
	}

	w := performRequest(t, router, "GET", "/projects/999", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("GetProjectByID() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestGetProjectByID_InvalidID(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/projects/:id", handler.GetProjectByID)

	w := performRequest(t, router, "GET", "/projects/invalid", nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("GetProjectByID() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestGetProjectByID_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/projects/:id", handler.GetProjectByID)

	mockRepo.getProjectByIDFunc = func(ctx context.Context, id int64) (*models.PortfolioProject, error) {
		return nil, errors.New("database error")
	}

	w := performRequest(t, router, "GET", "/projects/1", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetProjectByID() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

// =============================================================================
// Miniatures Handler Tests
// =============================================================================

func TestGetMiniatures_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/miniatures", handler.GetMiniatures)

	expectedMiniatures := []models.MiniatureProject{createTestMiniatureProject()}
	mockRepo.getAllMiniatureProjectsFunc = func(ctx context.Context) ([]models.MiniatureProject, error) {
		return expectedMiniatures, nil
	}

	w := performRequest(t, router, "GET", "/miniatures", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetMiniatures() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result []models.MiniatureProject
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("GetMiniatures() returned %d items, want 1", len(result))
	}
}

func TestGetMiniatures_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/miniatures", handler.GetMiniatures)

	mockRepo.getAllMiniatureProjectsFunc = func(ctx context.Context) ([]models.MiniatureProject, error) {
		return nil, errors.New("database error")
	}

	w := performRequest(t, router, "GET", "/miniatures", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetMiniatures() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

func TestGetMiniatureByID_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/miniatures/:id", handler.GetMiniatureByID)

	expectedMiniature := createTestMiniatureProject()
	mockRepo.getMiniatureProjectByIDFunc = func(ctx context.Context, id int64) (*models.MiniatureProject, error) {
		return &expectedMiniature, nil
	}

	w := performRequest(t, router, "GET", "/miniatures/1", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetMiniatureByID() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result models.MiniatureProject
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result.Title != testMiniatureName {
		t.Errorf("GetMiniatureByID() title = %s, want %s", result.Title, testMiniatureName)
	}
}

func TestGetMiniatureByID_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/miniatures/:id", handler.GetMiniatureByID)

	mockRepo.getMiniatureProjectByIDFunc = func(ctx context.Context, id int64) (*models.MiniatureProject, error) {
		return nil, gorm.ErrRecordNotFound
	}

	w := performRequest(t, router, "GET", "/miniatures/999", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("GetMiniatureByID() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestGetMiniatureByID_InvalidID(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/miniatures/:id", handler.GetMiniatureByID)

	w := performRequest(t, router, "GET", "/miniatures/invalid", nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("GetMiniatureByID() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestGetMiniatureByID_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/miniatures/:id", handler.GetMiniatureByID)

	mockRepo.getMiniatureProjectByIDFunc = func(ctx context.Context, id int64) (*models.MiniatureProject, error) {
		return nil, errors.New("database error")
	}

	w := performRequest(t, router, "GET", "/miniatures/1", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetMiniatureByID() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

// =============================================================================
// Miniature Themes Handler Tests
// =============================================================================

func TestGetMiniatureThemes_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/miniatures/themes", handler.GetMiniatureThemes)

	expectedThemes := []models.MiniatureTheme{createTestMiniatureTheme()}
	mockRepo.getAllMiniatureThemesFunc = func(ctx context.Context) ([]models.MiniatureTheme, error) {
		return expectedThemes, nil
	}

	w := performRequest(t, router, "GET", "/miniatures/themes", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetMiniatureThemes() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result []models.MiniatureTheme
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("GetMiniatureThemes() returned %d items, want 1", len(result))
	}
}

func TestGetMiniatureThemes_Empty(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/miniatures/themes", handler.GetMiniatureThemes)

	mockRepo.getAllMiniatureThemesFunc = func(ctx context.Context) ([]models.MiniatureTheme, error) {
		return []models.MiniatureTheme{}, nil
	}

	w := performRequest(t, router, "GET", "/miniatures/themes", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetMiniatureThemes() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetMiniatureThemes_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t)
	router.GET("/miniatures/themes", handler.GetMiniatureThemes)

	mockRepo.getAllMiniatureThemesFunc = func(ctx context.Context) ([]models.MiniatureTheme, error) {
		return nil, errors.New("database error")
	}

	w := performRequest(t, router, "GET", "/miniatures/themes", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetMiniatureThemes() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

// =============================================================================
// Context Propagation Tests
// =============================================================================

type ctxKey struct{}

func TestContextPropagation(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)

	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Add middleware that injects a sentinel value into the context
	router.Use(func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ctxKey{}, "test-marker")
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})

	router.GET("/certifications", handler.GetCertifications)

	var receivedCtx context.Context
	mockRepo.getAllCertificationsFunc = func(ctx context.Context) ([]models.Certification, error) {
		receivedCtx = ctx
		return []models.Certification{}, nil
	}

	w := performRequest(t, router, "GET", "/certifications", nil)

	if w.Code != http.StatusOK {
		t.Errorf("Request failed with status %d", w.Code)
	}

	if receivedCtx == nil {
		t.Error("Context was not propagated to repository")
	}

	// Verify the sentinel value was propagated through
	if receivedCtx.Value(ctxKey{}) != "test-marker" {
		t.Error("Context sentinel value was not propagated to repository")
	}
}

// =============================================================================
// Table-Driven Tests for ID Validation
// =============================================================================

func TestInvalidIDFormats(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t)

	router.GET("/projects/:id", handler.GetProjectByID)
	router.GET("/miniatures/:id", handler.GetMiniatureByID)

	// Note: Negative IDs are parseable by strconv.ParseInt, so they pass validation
	// and get a "not found" from the repository. Only non-numeric strings fail.
	tests := []struct {
		name      string
		path      string
		invalidID string
	}{
		{"project with string ID", "/projects/", "abc"},
		{"project with float ID", "/projects/", "1.5"},
		{"miniature with string ID", "/miniatures/", "xyz"},
		{"miniature with float ID", "/miniatures/", "3.14"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := performRequest(t, router, "GET", tt.path+tt.invalidID, nil)

			if w.Code != http.StatusBadRequest {
				t.Errorf("%s: status = %d, want %d", tt.name, w.Code, http.StatusBadRequest)
			}
		})
	}
}
