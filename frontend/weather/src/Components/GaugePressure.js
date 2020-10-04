import React, { Component } from 'react'  
import * as am4core from "@amcharts/amcharts4/core";
import * as am4charts from "@amcharts/amcharts4/charts";
//import am4themes_animated from "@amcharts/amcharts4/themes/animated";

/* Chart code */
// Themes begin
//am4core.useTheme(am4themes_animated);
// Themes end
class GaugePressure extends Component{
    
    componentDidMount() {  
        this.Gauge();
    }
    Gauge(){
        // create chart
        let chart = am4core.create("GaugePres", am4charts.GaugeChart);
        chart.innerRadius = am4core.percent(82);
        chart.numberFormatter.numberFormat = "#.";
        /**
         * Normal axis
         */
        let axis = chart.xAxes.push(new am4charts.ValueAxis());
        axis.title.text = "Pressão";
        axis.title.horizontalCenter = "middle";
        axis.min = 960;
        axis.max = 1060;
        axis.fontSize = 12;
        axis.strictMinMax = true;
        axis.renderer.radius = am4core.percent(102);
        axis.renderer.inside = false;
        axis.renderer.line.strokeOpacity = 1;
        axis.renderer.ticks.template.disabled = false;
        axis.renderer.ticks.template.strokeOpacity = 1;
        axis.renderer.ticks.template.length = 10;
        axis.renderer.grid.template.disabled = true;
        axis.renderer.labels.template.radius = 10;
        axis.renderer.labels.template.adapter.add("text", function(text) {
            return text;
        });

        /**
         * Axis for ranges
         */

        let colorSet = new am4core.ColorSet();

        let axis2 = chart.xAxes.push(new am4charts.ValueAxis());
        axis2.min = 960;
        axis2.max = 1060;
        axis2.strictMinMax = true;
        axis2.renderer.labels.template.disabled = true;
        axis2.renderer.ticks.template.disabled = true;
        axis2.renderer.grid.template.disabled = true;

        let range0 = axis2.axisRanges.create();
        range0.value = 960;
        range0.endValue = 1060;
        range0.axisFill.fillOpacity = 1;
        range0.axisFill.fill = colorSet.getIndex(0);

        let range1 = axis2.axisRanges.create();
        range1.value = 960;
        range1.endValue = 1060;
        range1.axisFill.fillOpacity = 1;
        range1.axisFill.fill = colorSet.getIndex(9);

        /**
         * Label
         */

        let label = chart.radarContainer.createChild(am4core.Label);
        label.isMeasured = true;
        label.fontSize = 25;
        label.x = am4core.percent(50);
        label.y = am4core.percent(100);
        label.horizontalCenter = "middle";
        label.verticalCenter = "bottom";
        label.text = "0";

        let labelText = chart.radarContainer.createChild(am4core.Label);
        labelText.isMeasured = false;
        labelText.fontSize = 20;
        labelText.x = am4core.percent(100);
        labelText.y = -30;
        labelText.horizontalCenter = "middle";
        labelText.verticalCenter = "bottom";
        labelText.text = "hpa";
        

        /**
         * Hand
         */

        let hand = chart.hands.push(new am4charts.ClockHand());
        hand.axis = axis2;
        hand.innerRadius = am4core.percent(40);
        hand.startWidth = 10;
        hand.pin.disabled = true;
        hand.value = 1010;

        hand.events.on("propertychanged", function(ev) {
            range0.endValue = ev.target.value;
            range1.value = ev.target.value;
            label.text = axis2.positionToValue(hand.currentPosition).toFixed(3);
            axis2.invalidate();
        });

        setInterval(() =>{
            let value = this.props.iotdata.Pres;
            new am4core.Animation(hand, {
                property: "value",
                to: value
            }, 1000, am4core.ease.cubicOut).start();
        }, 2000);
    }
    componentWillUnmount() {  
        if (this.chart) {  
            this.chart.dispose();  
        }  
    }
    
    render() {  
        return (  
            <div id="GaugePres" className="gauge"></div>
        )  
    }  

}

export default GaugePressure;