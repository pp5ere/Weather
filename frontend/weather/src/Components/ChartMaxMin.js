import React, { Component } from 'react'  
import * as am4core from "@amcharts/amcharts4/core";  
import * as am4charts from "@amcharts/amcharts4/charts";  
//import am4themes_animated from "@amcharts/amcharts4/themes/animated";  
import am4lang_pt_BR from "@amcharts/amcharts4/lang/pt_BR";
  
//am4core.useTheme(am4themes_animated);

class ChartMaxMin extends Component {  
  
    componentDidMount() {  
        this.Chart(this.props.weather, "C°");
    }  

    componentDidUpdate(oldProps){
        if (oldProps.weather !== this.props.weather){
            this.chart.data = this.props.weather;
        }
    }

    Chart(weather, unit){
        let chart = am4core.create("MaxMinChart", am4charts.XYChart);  
        chart.language.locale = am4lang_pt_BR;
        
        // Add data  
        chart.data = weather;
        let dateAxis = chart.xAxes.push(new am4charts.DateAxis());
        
        dateAxis.tooltipDateFormat = "d MMMM YYYY";
        dateAxis.renderer.grid.template.location = 0;  
        dateAxis.renderer.minGridDistance = 80;
        chart.colors.list = [
            am4core.color("#8c93ba"),
            am4core.color("#a41e1e"),            
        ];

        let valueAxis = chart.yAxes.push(new am4charts.ValueAxis());
        valueAxis.tooltip.disabled = true;
        valueAxis.title.text = unit;

        let series = chart.series.push(new am4charts.LineSeries());
        series.dataFields.dateX = "Data";
        series.dataFields.valueY = "MinTempC";
        series.yAxis = valueAxis;
        series.name = "Temperatura Mínima por Dia";
        series.tooltipText = '{name}\n[bold font-size: 20]{valueY} '+unit+'[/]';

        let seriesMax = chart.series.push(new am4charts.LineSeries());
        seriesMax.dataFields.dateX = "Data";
        seriesMax.dataFields.valueY = "MaxTempC";
        seriesMax.yAxis = valueAxis;
        seriesMax.name = "Temperatura Máxima por Dia";
        seriesMax.tooltipText = '{name}\n[bold font-size: 20]{valueY} '+unit+'[/]';
        //series.fillOpacity = 0.3;
        
        /*let bullet3 = series.bullets.push(new am4charts.CircleBullet());  
        bullet3.circle.radius = 3;*/  
        /*bullet3.circle.strokeWidth = 2;  
        bullet3.circle.fill = am4core.color("#fff");*/

        chart.cursor = new am4charts.XYCursor();
        chart.cursor.lineY.opacity = 0;
        /*chart.scrollbarX = new am4charts.XYChartScrollbar();
        chart.scrollbarX.fontSize = 10;
        chart.scrollbarX.contentHeight = 1;
        chart.scrollbarX.series.push(series);
        chart.scrollbarX.parent = chart.bottomAxesContainer;*/
        chart.scrollbarX = new am4core.Scrollbar();            
        chart.scrollbarX.parent = chart.bottomAxesContainer;
        
        // Add legend  
        chart.legend = new am4charts.Legend();  
        chart.legend.position = "top";  

        dateAxis.start = 0.5;
        dateAxis.keepSelection = true;
        
        this.chart = chart;  
    }
  
    componentWillUnmount() {  
        if (this.chart) {  
            this.chart.dispose();  
        }  
    }  
    render() {  
        return (  
            <div id="MaxMinChart" className="chartFull"></div>
        )  
    }  
}  

export default ChartMaxMin;