
function buildHeatmap(graphData, domElement) { 
        var tdy = new Date();
        var cal = new CalHeatMap();
        cal.init({
            itemName: "commit",
            considerMissingDataAsZero: true,
            cellSize: cellSize,
            legendCellSize: 15,
            legend: [0, 4, 10, 20, 30], 
            nimationDuration: 600,
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
