var parseNameAndImgFromGit = function(url) {
    var img = document.createElement("img");
    img.className = "circle";
    var name;
    if ( url.indexOf("github.com") > -1 ) { 
        img.src = "/static/images/github.png";
        name = parseGitName(url);
    } else if ( url.indexOf("hub.jazz.net") > -1 ) {
        img.src = "/static/images/bluemix.png";
        name = parseJazzName(url);
    } else if ( url.indexOf("github.rtp") > -1 ) { 
        img.src = "/static/images/gitlab.png";
        name = parseGitlabName(url);
    } else if ( url.indexOf("github.ibm.com") > -1 ) { 
        img.src = "/static/images/github.png";
        name = parseGitName(url);
    }
    return {
        img: img,
        name: name
    };
};
var parseNameFromGit = function(url) {
    var name;
    if ( url.indexOf("github.com") > -1 ) { 
        name = parseGitName(url);
    } else if ( url.indexOf("hub.jazz.net") > -1 ) {
        name = parseJazzName(url);
    } else if ( url.indexOf("github.rtp") > -1 ) { 
        name = parseGitlabName(url);
    } else if ( url.indexOf("github.ibm.com") > -1 ) { 
        name = parseGitName(url);
    }
    return name;
};
var parseGitName = function(url) {
    var strs = url.split("/");
    return strs[strs.length-1].replace(".git", "");
};

var parseJazzName = function(url) { 
    var strs = url.split("/");
    return strs[strs.length-1];
};

var parseGitlabName = function(url) {
    var strs = url.split("/");
    return strs[strs.length-1].replace(".git", "");
};
