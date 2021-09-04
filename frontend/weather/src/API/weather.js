import axios from 'axios';
//import configJson from './config.json';

/*const DOMAIN = "http://pp5ere.sytes.net";
const HOST = DOMAIN + ":9000";*/
const HOST = window.location.protocol + "//" + window.location.hostname + ":" + window.location.port;//configJson.APIPort;
const dateToString = d => `${ d.getFullYear() }-${('00' + (d.getMonth() + 1)).slice(-2)}-${('00' + d.getDate()).slice(-2)}`;

const getWeatherByDate = async (date) =>{
    if (!date) {
        console.log('Data: '+date);
        return [];        
    } else {
        try {
            const res = await axios.get(`${HOST}/weather/` + dateToString(date));
            setDateArray(res.data.Weather);
            return res.data.Weather;
        } catch (err) {
            console.error(err);
            return [];
        }
    }
}

const setDateArray = (array) => {
    if (array != null){
        array.forEach((item)=> {
          item.Data = new Date(item.Data);
      });
    }
    return array;
}

const getAllWeather = async () => {
    try {
        const res = await axios.get(`${HOST}/weather`);
        //console.log('res: ' + res.data.Weather);
        setDateArray(res.data.Weather);
        return res.data.Weather;
    } catch (err) {
        console.error(err);
        return [];
    }
}

const getMaxMinTempCPerDay = async (date) => {
    if (!date) {
        console.log('Data: ' + date);
        return [];        
    } else {
        try {
            const res = await axios.get(`${HOST}/maxmintemp/` + dateToString(date));
            setDateArray(res.data.Weather);
            //console.log('res: ' + res.data.Weather);
            return res.data.Weather;
        } catch (err) {
            console.error(err);
            return [];
        }
    }
}

const getDataFromSensor = async() =>{
    try {
        const res = await axios.get(`${HOST}/iotdata`);
        return res.data;
    } catch (error) {
        console.log(error);
        return {};
    }
}

export {
    getWeatherByDate, 
    getAllWeather,
    getMaxMinTempCPerDay,
    dateToString,
    getDataFromSensor,
}
