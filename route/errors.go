package route

const (
	ERROR_REPO_URL = "An error occured while parsing you git URL. Ensure that you have entered the URL correctly, and that you are using an HTTP or HTTPS URL (ie not the git@* style URL which requires the use of SSH). Also recall that only IBM Jazz and GitHub URLs are currently only supported. Other provider might still work. Please contact me at <a href='mailto:schurman@ca.ibm.com'>schurman@ca.ibm.com</a> for more help."

	ERROR_CLONE_REPO = "An error occured while cloning your repository using the git CLI. Ensure that the URL for your repository is entered correcrtly, and that you are using an HTTP or HTTPS URL (ie not the git@* style URL which requires the use of SSH). Also ensure that the repo is public, or that you have first added the Git-Monitor Bot as a member if it is private. If you believe that all this criteria is met, then please try your request again, as this error could be random fluke due to connection loss while downloading."

	ERROR_PARSE_PAGE = "An error occured rendering the requested page. Please try again. If the problem persists, please contact me at schurman@ca.ibm.com"
)
