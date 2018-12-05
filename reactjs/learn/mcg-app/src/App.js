import React from 'react';
//import PropTypes from 'prop-types';
import axios from 'axios';
import logo from './logo.svg';
import './App.css';
import './RoyalLog.js'
import { Base, AppInfo, ErrorInfo, ErrorEvent, ServiceMetricInfo, ServiceMetricEvent } from './RoyalLog.js';

// Helper for logging
class Logger {
  static factoryBaseError(traceId, vdsId) {
    const base = new Base("Error", traceId, vdsId, 
      new Date().toISOString(), 'profile')
    return base
  }
  static factoryApp() {
    let ai = new AppInfo("ReactDemo", "web", "1.0", "chrome", "GCP")
    return ai
  }

  static factoryError(errId, errMsg, context) {
    let e = new ErrorInfo(errId, errMsg, "", context, 
      [], "", true, "")
    return e
  }

  static remoteLog(logEvent) {
    let local = "http://192.168.64.26:31380/royal/api/logger "
    let remote = "http://35.245.49.124/royal/api/logger"
    let url = remote 

    fetch(url, {
      method: 'post', 
      body: JSON.stringify(logEvent)
    })
    .then( (res) => res.json())
    .then( (json) => {
        console.log("remote log resp => ", json)     
    })
    .catch((err)=>console.log(err))
  }

}


class ReservationForm extends React.Component {
  constructor() {
    super()
    this.state = {
      loaded: false,
      errorMsg: "",
      profile: { }
      //reservations: []
    }
  }

  fetchData(vdsId) {
    // example vdsId = bjm100
    /* uncomment for hit lower level booking service
    let uri = "http://35.245.49.124/royal/api/bookings/"
    */
    let uri = "http://35.245.49.124/royal/api/profile/"
    let url = uri + vdsId
    let start = Date.now()
    axios({
      url: url,
      method: 'get',
    })
    .then( res => {
        //const bookingsArray =  res.data 
        //this.setState({reservations: bookingsArray})
        //console.log(bookingsArray)
        const profile = res.data
        this.setState({profile: profile})
        this.setState({loaded: true})

        console.log(profile)

        var latency = Date.now() - start;
        let metric = new ServiceMetricInfo(
          "booking", "bookings", "GET", latency, 
          null, { vdsId: vdsId}
        )

        let logEvent = new ServiceMetricEvent(
          Logger.factoryBaseError( "abcd120", vdsId),
          Logger.factoryApp(),
          metric)
        
        // send log event to server for Splunk
        Logger.remoteLog(logEvent)
        
        

    }, error => {
      // Errors, like not found, etc.
      this.setState({errorMsg: error.message})
      console.log(error);
      var b = Logger.factoryBaseError( "abcd120", "bjm100")
      var ai = Logger.factoryApp()
      var e = Logger.factoryError("ErrProfileFetch", "Cant load profile", {} )
      var ee = new ErrorEvent(b, ai, e)
      console.log(ee)
      Logger.remoteLog(ee)
    });
  }

  handleSubmit = (evt) => {
    evt.preventDefault() 

    let vdsId = this.inputField.value

    // clear current state of reservation info
    //this.setState({loaded: false, errorMsg: "", reservations:  []})
    this.setState({loaded: false, errorMsg: "", profile:  {} })
  

    console.log("fetching reservation for: ", vdsId, "...")
    this.fetchData(vdsId)
  }

  render() {

    // render empty div tag if no reservations are currently loaded.
    var rezes = <div />

    if (this.state.errorMsg.length > 0) {
      rezes = <div>{this.state.errorMsg}</div>
    }
  
    let p = this.state.profile
    if (this.state.loaded) {
        var reservations = p.reservations.map( (rez) => {
            let sail_date = rez.sailDate.substring(0,10)
            return(<tr key={rez.shipCode}><td>{rez.shipCode}</td><td align="left">{rez.shipName} </td><td> {sail_date} </td></tr>)
        })
        rezes = <table className="profile"><tbody>
          <tr key="1"><td colSpan="3" className="nameRow"> {p.firstName} {p.lastName}</td></tr>
          {reservations}
        </tbody></table>

    } 


    return (
      <div>
      <form onSubmit={this.handleSubmit}>
        <label>Enter VdsId: </label>
        <input type="text" ref={node => this.inputField = node}></input>
        <button type="submit">Find Reservations</button>
      </form>
      <hr />
        <div>
         {rezes}
        </div>
      </div>

    )
  }
}


// Old school component style
class App extends React.Component {
  constructor() {
    super();
    this.state = {
      msg: 'Reservations'
    }
  }

  update(evt) {
    let value = evt.target.value
    this.setState({msg: value})

  }

  render() {
    return (
      <div className="App">
      <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
      </header>
      <div>
        <h1>Demo SPA App - Sailing Reservation Finder</h1>
        < ReservationForm />
      </div>

      </div>
    )
  }
  
}





export default App;
