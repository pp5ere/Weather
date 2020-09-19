    import React, { Component } from 'react'  
    import * as am4core from "@amcharts/amcharts4/core";  
    import * as am4charts from "@amcharts/amcharts4/charts";  
    //import am4themes_animated from "@amcharts/amcharts4/themes/animated";  
    import am4lang_pt_BR from "@amcharts/amcharts4/lang/pt_BR";
    
      
    //am4core.useTheme(am4themes_animated);

    class ChartTemperature extends Component {  
      
        componentDidMount() {  
            this.Chart(this.props.weather, "C°");
            //this.Chart("C°");
        }  
        
        componentDidUpdate(oldProps){
            if (oldProps.weather !== this.props.weather){
                this.chart.data = this.props.weather;
            }
           /* this.chart.dataSource.load();
            this.chart.dataSource.events.on("done", function(ev) {
                // Data loaded and parsed
                //console.log(ev.target.data.Weather);
                let weatherArray = ev.target.data.Weather;
                if (weatherArray != null){
                weatherArray.forEach((item)=> {
                    item.Data = new Date(item.Data);
                });
                this.chart.data =weatherArray;
                console.log(weatherArray);
            } 
            });*/
            //console.log(this.props.weather);
            /*if (oldProps.date !== this.props.date){
                this.chart.dataSource.url = "http://pp5ere.sytes.net:9000/weather/" +  dateToString(this.props.date);
                console.log(this.chart.dataSource.url);
                this.chart.dataSource.load();
                if (this.props.date.toISOString().slice(0,10)  === new Date().toISOString().slice(0,10)){
                    this.chart.dataSource.incremental = true;
                    this.chart.dataSource.reloadFrequency = 60000;
                }else {
                    this.chart.dataSource.incremental = false;
                    this.chart.dataSource.reloadFrequency = 0;
                }
            }*/
        }

        Chart(weather,unit){
            let chart = am4core.create("TempChart", am4charts.XYChart);  
            //let dataSource = new am4core.DataSource();
            chart.language.locale = am4lang_pt_BR;
            //this.loadData(chart);
            
            // Add data  
            chart.data = weather;
                        
            //chart.data = chart.data.Weather;
            /*chart.dataSource.url = "http://pp5ere.sytes.net:9000/weather/" + dateToString(this.props.date);
            chart.dataSource.parser = am4core.JSONParser;
            chart.dataSource.load();
            chart.dataSource.events.on("done", function(ev) {
                // Data loaded and parsed
                console.log(ev.target.data.Weather);
            });
            chart.dataSource.load();
            chart.dataSource.events.on("done", function(ev) {
                // Data loaded and parsed
                //console.log(ev.target.data.Weather);
                let weatherArray = ev.target.data.Weather;
                if (weatherArray != null){
                weatherArray.forEach((item)=> {
                    item.Data = new Date(item.Data);
                });
                chart.data =weatherArray;
            } 
            });
            
            chart.dataSource.incremental = true;
            chart.dataSource.reloadFrequency = 60000;*/
            
            let dateAxis = chart.xAxes.push(new am4charts.DateAxis());
            
            dateAxis.tooltipDateFormat = "HH:mm, d MMMM";
            dateAxis.renderer.grid.template.location = 0;  
            dateAxis.renderer.minGridDistance = 50;
            chart.colors.list = [
                am4core.color("#FF9E01"),
                am4core.color("#a41e1e"),
                am4core.color("#237d2a"),
            ];

            let valueAxis = chart.yAxes.push(new am4charts.ValueAxis());
            valueAxis.tooltip.disabled = true;
            valueAxis.title.text = unit;

            let series = chart.series.push(new am4charts.LineSeries());
            series.dataFields.dateX = "Data";
            series.dataFields.valueY = "TempC";
            series.yAxis = valueAxis;
            series.name = "Temperatura";
            series.tooltipText = '{name}\n[bold font-size: 20]{valueY} '+unit+'[/]';
            //series.fillOpacity = 0.3;
            
            /*let bullet3 = series.bullets.push(new am4charts.CircleBullet());  
            bullet3.circle.radius = 3;*/  
            /*bullet3.circle.strokeWidth = 2;  
            bullet3.circle.fill = am4core.color("#fff");*/

            let seriesTr = chart.series.push(new am4charts.LineSeries());
            seriesTr.dataFields.dateX = "Data";
            seriesTr.dataFields.valueY = "Hi";
            seriesTr.yAxis = valueAxis;
            seriesTr.name = "Sensação";
            seriesTr.tooltipText = '{name}\n[bold font-size: 20]{valueY} '+unit+'[/]';
            //seriesTr.fillOpacity = 0.3;

            /*let bulletTr = seriesTr.bullets.push(new am4charts.CircleBullet());  
            bulletTr.circle.radius = 3; */ 

            let seriesDp = chart.series.push(new am4charts.LineSeries());
            seriesDp.dataFields.dateX = "Data";
            seriesDp.dataFields.valueY = "DewPoint";
            seriesDp.yAxis = valueAxis;
            seriesDp.name = "Ponto de orvalho";
            seriesDp.tooltipText = '{name}\n[bold font-size: 20]{valueY} '+unit+'[/]';
            //seriesDp.fillOpacity = 0.3;

            /*let bulletDp = seriesDp.bullets.push(new am4charts.CircleBullet());  
            bulletDp.circle.radius = 3; */ 

            chart.cursor = new am4charts.XYCursor();
            chart.cursor.lineY.opacity = 0;
            /*chart.scrollbarX = new am4charts.XYChartScrollbar();
            chart.scrollbarX.fontSize = 10;
            chart.scrollbarX.contentHeight = 1;
            chart.scrollbarX.series.push(series);
            chart.scrollbarX.parent = chart.bottomAxesContainer;*/   

            chart.scrollbarX = new am4core.Scrollbar();            
            chart.scrollbarX.parent = chart.bottomAxesContainer;
            /*chart.scrollbarX.startGrip.hide();
            chart.scrollbarX.endGrip.hide();
            chart.scrollbarX.start = 0;
            chart.scrollbarX.end = 0.25;
            chart.zoomOutButton = new am4core.ZoomOutButton();
            chart.zoomOutButton.hide();*/
              
            
            // Add legend  
            chart.legend = new am4charts.Legend();  
            chart.legend.position = "top";  

            dateAxis.start = 0; //inicia com zoom total 
            dateAxis.keepSelection = true;
            
            this.chart = chart;  
        }

        componentWillUnmount() {  
            if (this.chart) {  
                this.chart.dispose(); 
                delete this.chart; 
            }  
        }  

        /*loadData(chart){
            chart.dataSource.url = "http://pp5ere.sytes.net:9000/weather/" + dateToString(this.props.date);            
            chart.dataSource.load();
            chart.dataSource.events.on("done", function(ev) {
                // Data loaded and parsed
                let weatherArray = ev.target.data.Weather;
                if (weatherArray != null){
                weatherArray.forEach((item)=> {
                    item.Data = new Date(item.Data);
                });
                chart.data =weatherArray;
            } 
            });
            
            
            chart.dataSource.incremental = true;
            chart.dataSource.reloadFrequency = 60000;
        }*/
      

        render() {  
            return (  
                <div id="TempChart" className="chartFull"></div>
            )  
        }  
    }  

export default ChartTemperature;