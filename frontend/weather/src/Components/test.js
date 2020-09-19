import React, { Component } from 'react'  
import * as am4core from "@amcharts/amcharts4/core";  
import * as am4charts from "@amcharts/amcharts4/charts";  


class Test extends Component {  
    componentDidMount() {  
        this.charttest();
    }  
    charttest(){

let chart = am4core.create("chartdiv", am4charts.XYChart);

// Set up data source
//chart.dataSource.url ="https://s3-us-west-2.amazonaws.com/s.cdpn.io/t-160/sample_data_serial.json";
chart.dataSource.url ="http://pp5ere.sytes.net:9000/weather/2020-04-29";
        
            chart.dataSource.parser = new am4core.JSONParser();
            chart.dataSource.parser.options.emptyAs = 0;
            chart.dataSource.incremental = true;
            chart.dataSource.reloadFrequency = 5000;
            var t = chart.dataSource.data.Weather;
            console.log(t);

// Create axes
var categoryAxis = chart.xAxes.push(new am4charts.CategoryAxis());
categoryAxis.dataFields.category = "year";

// Create value axis
var valueAxis = chart.yAxes.push(new am4charts.ValueAxis());

// Create series
var series1 = chart.series.push(new am4charts.LineSeries());
series1.dataFields.valueY = "cars";
series1.dataFields.categoryX = "year";
series1.name = "Cars";
series1.strokeWidth = 3;
series1.tensionX = 0.7;
series1.bullets.push(new am4charts.CircleBullet());

var series2 = chart.series.push(new am4charts.LineSeries());
series2.dataFields.valueY = "motorcycles";
series2.dataFields.categoryX = "year";
series2.name = "Motorcycles";
series2.strokeWidth = 3;
series2.tensionX = 0.7;
series2.bullets.push(new am4charts.CircleBullet());

var series3 = chart.series.push(new am4charts.LineSeries());
series3.dataFields.valueY = "bicycles";
series3.dataFields.categoryX = "year";
series3.name = "Bicycles";
series3.strokeWidth = 3;
series3.tensionX = 0.7;
series3.bullets.push(new am4charts.CircleBullet());

// Add legend
chart.legend = new am4charts.Legend();
this.chart = chart; 
    }
    render() {  
        return (  
            <div id="chartdiv" className="chartFull"></div>
        )  
    }  
}

export default Test;