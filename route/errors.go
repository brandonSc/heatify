package route

const (
	ERROR_REPO_URL = "An error occured while parsing you git URL. Ensure that you have entered the URL correctly, and that you are using an HTTP or HTTPS URL (ie not the git@* style URL which requires the use of SSH). Also recall that only IBM Jazz and GitHub URLs are currently only supported. Other provider might still work. Please contact me at <a href='mailto:schurman@ca.ibm.com'>schurman@ca.ibm.com</a> for assistance."

	ERROR_CLONE_REPO = "Git-Monitor encountered an error while inspecting your repository. Ensure that the URL for your repository is entered correcrtly, and that you are using an HTTP or HTTPS URL (ie not the 'git@' prefixed URL which requires the use of SSH). Also be sure that the repo is public, or that you have first added the Git-Monitor Bot as a member if it is private. If you believe that this criteria is met, then please try your request again, as this could be a random error due to something like connection loss while downloading the repository data."

	ERROR_PARSE_PAGE = "An error occured rendering the requested page. Please try again. If the problem persists, contact me at schurman@ca.ibm.com"
)
