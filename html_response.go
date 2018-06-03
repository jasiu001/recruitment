package main

type InfoPage struct {
	Name         string
	Email        string
	Repositories []PublicRepository
}

type PublicRepository struct {
	RepoName  string
	Languages map[string]string
}

type ErrorPage struct {
	ErrorInfo string
}

func NewInfoPage(account *Account) InfoPage {
	repositories := []PublicRepository{}
	for _, repo := range account.GetRepositories() {
		lang, err := account.GetRepositoryLanguage(repo)
		if err != nil {
			lang = map[string]string{"undefined": "0"}
		}
		repositories = append(repositories, PublicRepository{
			RepoName:  repo,
			Languages: lang,
		})
	}

	return InfoPage{
		Name:         account.UserName(),
		Email:        account.GetEmail(),
		Repositories: repositories,
	}
}

func NewErrorPage(info string) ErrorPage {
	return ErrorPage{
		ErrorInfo: info,
	}
}
