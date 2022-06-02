import React, { Component } from 'react'  
import ChartTemperature from './ChartTemperature';
import ChartHum from './ChartHum';
import ChartPres from './ChartPres';
import ChartMaxMin from './ChartMaxMin';
import { getWeatherByDate } from '../API/weather'
import Calendar from 'react-calendar';
import 'react-calendar/dist/Calendar.css';
import GaugesContainer from './GaugeContainer';

export default class ChartContainer extends Component{
    constructor (props){
        super(props);
        this.state = {
            date: new Date(),
            weather: [],
        };
    }
    
    componentDidMount() {  
        this.getWeather(this.state.date);
        this.timeValue = this.startTimer();
    }  

    componentWillUnmount() {
        this.stopTimer(this.timeValue);
    }

    async getWeather(d){
        this.setState({date: d, weather: await getWeatherByDate(d)});
    }

    startTimer(){
        this.isTimerRunning = true;
        return setInterval(this.recallWeather, 180000);
    }

    stopTimer(idTimer){
        this.isTimerRunning = false;
        return clearInterval(idTimer);
    }

    recallWeather = () =>{        
        //console.log("RECALL: CurrentDate: "+this.getYearMouthDayFromDate(new Date()) +" Date: "+this.getYearMouthDayFromDate(this.state.date));
        const newDate = new Date();
        if (getYearMouthDayFromDate(newDate) > getYearMouthDayFromDate(this.state.date)){
            //console.log("New Day");
            this.getWeather(newDate);
        }else {
            this.getWeather(this.state.date);
        }
    }

    onChange = (d) => {         
        this.getWeather(d);
        //console.log("CurrentDate: "+this.getYearMouthDayFromDate(new Date()) +" Date: "+this.getYearMouthDayFromDate(d));
        if (getYearMouthDayFromDate(new Date()) < getYearMouthDayFromDate(d) || getYearMouthDayFromDate(new Date()) > getYearMouthDayFromDate(d)){
            if (this.isTimerRunning){
                this.stopTimer(this.timeValue);
                //console.log("Clear: "+this.timeValue);
            }
            
        }else if (getYearMouthDayFromDate(new Date()) === getYearMouthDayFromDate(d)){
            if (!this.isTimerRunning){
                this.timeValue = this.startTimer();
            }
        }
    };
    

    render(){
        return (
            <div>
                <div className="topContainer"> 
                    <div className="topDate">                        
                        <div className="date">
                            <div className="text">{dateToStringBr(this.state.date)}</div>
                            <div className="calendar">
                                <Calendar onChange={this.onChange} value={this.state.date}/>
                            </div>                        
                        </div>                        
                    </div>
                    
                    <div>                        
                        <GaugesContainer/>                                                    
                        <ChartTemperature weather={this.state.weather}/>
                        <div className="chartContainer"> 
                            <ChartHum weather={this.state.weather}/>
                            <ChartPres weather={this.state.weather}/>
                        </div>
                        <ChartMaxMin date={this.state.date}/>
                    </div>
                       
                </div>
            </div>
        );
    }
}

const dateToStringBr = d => `${('00' + d.getDate()).slice(-2)}/${('00' + (d.getMonth() + 1)).slice(-2)}/${ d.getFullYear() }`;
const getYearMouthDayFromDate = (d1)=>{
    return d1.getFullYear()+d1.getMonth()+d1.getDate();
}