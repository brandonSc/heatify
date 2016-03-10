
/**
 * build a general heatmap of the graphData and attach to the the domElement
 */
function buildHeatmap(graphData, domElement) { 
    var tdy = new Date();
    var cal = new CalHeatMap();
    cal.init({
        itemName: "commit",
        considerMissingDataAsZero: true,
        cellSize: cellSize,
        legendCellSize: 15,
        legend: [0, 4, 10, 20, 30], 
        animationDuration: 600,
        range: 12,
        domain: "month",
        subDomain: "day",
        tooltip: true,
        start: new Date(tdy.getFullYear(), tdy.getMonth()-11, tdy.getDay()),
        //start:  new Date(startDate.getFullYear(), startDate.getMonth(), startDate.getDay()),
        data: graphData,
    });

    $("#left-pan").on("click", function(e) {  
        e.preventDefault();
        cal.next();
    });

    $("#right-pan").on("click", function(e) {
        e.preventDefault();
        cal.previous();
    });
}

/**
 * build a heatmap for a squad member's data and attach it to the given element
 */
function buildSquadMemberHeatmap(member, squad, rank, element, cellSize, terminate) {
    var memberId = memberNameToId(trimGitAuthorToMember(member));
    var url = "/api/commits/squad/user?user="+member+"&squad="+squad;
    $.get(url, function(response) {
        var memberElem = document.getElementById("heatmap-"+memberId);
        var data = commitsToCalData(JSON.parse(response));

        var elem = document.createElement("div");
        elem.style = "padding-top: 15px; padding-bottom:50px;";
        

        leftPan = document.createElement("a");
        leftPan.id = "left-pan-"+memberId;
        leftPan.className = "waves-effect waves-teal btn-flat";
        leftPan.style = "font-size: 24px; left: 0%";
        leftPan.innerHTML = "&#62;";

        rightPan = document.createElement("a");
        rightPan.id = "right-pan-"+memberId;
        rightPan.className = "waves-effect waves-teal btn-flat";
        rightPan.style = "font-size: 24px; right: 0%";
        rightPan.innerHTML = "&#60;";

        var memElem = document.createElement("h5");
        memElem.innerHTML = trimGitAuthorToMember(member) + ' &nbsp; #' + rank;
    //    var rankElem = document.createElement("h6");
    //    rankElem.innerHTML = '#'+rank;

        element.appendChild(memElem);
        element.appendChild(rightPan);
        element.appendChild(leftPan);
        //element.appendChild(rankElem);
        element.appendChild(elem);

        if ( cellSize === null || cellSize === 0 ) 
            cellSize = 15;

        var tdy = new Date();
        var cal = new CalHeatMap();
        cal.init({
            itemSelector: elem,
            itemName: "commit",
            considerMissingDataAsZero: true,
            cellSize: cellSize,
            legend: [0, 4, 10, 20, 30], 
            animationDuration: 600,
            range: 10,
            domain: "month",
            subDomain: "day",
            tooltip: true,
            displayLegend: false,
            start: new Date(tdy.getFullYear(), tdy.getMonth()-11, tdy.getDay()),
            data: data,
        });
        
        $("#left-pan-"+memberId).on("click", function(e) {  
            e.preventDefault();
            cal.next();
        });

        $("#right-pan-"+memberId).on("click", function(e) {
            e.preventDefault();
            cal.previous();
        });

        if ( terminate ) { 
            var target = document.getElementById('spinner')
            target.innerHTML = "";
        }

    });
}

/**
 * given a list of members (git authors) 
 * query the Heatify API for their commit data
 * and build a heatmap for each
 */
function buildAllMembersHeatmaps(squad, element, cellSize) { 
    var members = squad.members;
    // pre add elements for member heatmaps
    for ( var i=0; i<members.length; i++ ) { 
        var member = memberNameToId(trimGitAuthorToMember(members[i]));
        var memberElem = document.createElement("heatmap-"+member);
        memberElem.id = "heatmap-"+member;
        memberElem.style = "margin: auto; display: inline-block;";
        element.appendChild(memberElem);
    }



    // do this in a separate loop so elements are not added as data is loaded
    for ( var j=0; j<members.length; j++ ) { 
        var member = memberNameToId(trimGitAuthorToMember(members[j]));
        var terminate = false;
        if ( j == members.length-1 ) 
            terminate = true;
            
        buildSquadMemberHeatmap(members[j], squad.name, j+1, document.getElementById('heatmap-'+member), cellSize, terminate);
    }
}

/**
 * extract the members name from the git author
 * e.g. :
 * consumes: Brandon Schurman <schurman@ca.ibm.com>
 * produces: Brandon Schurman
 */
function trimGitAuthorToMember(author) {
    var i = author.indexOf(" <");
    var j = author.indexOf(">");
    if ( i < j ) { 
        var member = author.substring(0, i);
        return member;
    } else { 
        console.log("error parsing member name from git author in heatmapBuilder.js");
        return author;
    }
}

function memberNameToId(member) { 
    return member.replace(" ", "-");
}

function commitsToCalData(commits) { 
    var graphData = {};
    var startDate = new Date();
    var endDate  = new Date();
    for ( var i=0; i<commits.length; i++ ) { 
        var d = new Date(commits[i].date);
        d.setDate(d.getDate() + 1);
        graphData[d.getTime()/1000] = commits[i].commits;
        if ( d.getTime() < startDate.getTime() ) 
            startDate = d
        if ( d.getTime() > endDate.getTime() ) 
            endDate = d;
        totalCommits = totalCommits + commits[i].commits;
    }
    return graphData;
}
