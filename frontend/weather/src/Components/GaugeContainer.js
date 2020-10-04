import React, { Component } from 'react'  
import {getDataFromSensor} from '../API/weather'
import 'react-calendar/dist/Calendar.css';
import GaugeHumidity from './GaugeHumidity';
import GaugePressure from './GaugePressure';
import GaugeTemperature from './GaugeTemperature';

export default class GaugesContainer extends Component{
    constructor(props) {
        super(props);
        this.state = {
            iotdata:{}
        };
      }
    
    componentDidMount() {
        this.getDataFromIot();        
        this.timerID = setInterval(
            () => this.getDataFromIot(),
            3000
          );
    }
    
    componentWillUnmount() {
        clearInterval(this.timerID);
    }
    
    async getDataFromIot() {          
        this.setState(await getDataFromSensor());                
    }
    
    render(){
        return (                               
            <div className="chartGauge"> 
                <GaugeTemperature iotdata={this.state}/>
                <GaugeHumidity iotdata={this.state}/>
                <GaugePressure iotdata={this.state}/>
            </div>                                           
        );
    }
}

