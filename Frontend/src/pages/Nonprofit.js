import React, { Component } from 'react'
import {Button} from 'react-bootstrap'
import axios from 'axios'
import styled from 'styled-components'

const StyledCol = styled.td`
    border: 1px solid black;
    padding-right: 20px;
    padding-left: 20px;
    padding-top: 5px;
    padding-bottom: 5px;
`;

const StyledHead = styled.th`
    border: 1px solid black;
    padding-right: 20px;
    padding-left: 20px;
    padding-top: 5px;
    padding-bottom: 5px;
`;

const api = axios.create({
    baseURL: 'http://localhost:8080/'
  })

export class Nonprofit extends Component {
    state = {
        events: []
    }

    constructor() {
        super();
        api.get('/Events').then(res => {
            console.log(res.data)
            this.setState({events : res.data})
        })
    }

    render(){
        return(
            <div>
                <div class='row'><img src="http://placeimg.com/150/150/animals" rounded></img>
                <div class = 'col'><h1>Indianapolis ASPCA:</h1><p>Our mission is to care for Indianapolis' displaced animals while find loving homes for each of them! 
                    Please see below for oppurtunities to help out with our events, we could always use more volunteers! Our building is located at 3400 Northwood Lane,
                    Indianapolis, IN, 46268!</p></div></div>
                <div class='row'><h3>Upcoming Events:</h3></div>
                <tr>
                    <StyledHead>Event</StyledHead>
                    <StyledHead>Date</StyledHead>
                    <StyledHead>View Event</StyledHead>
                </tr>
                {this.state.events.map(event => 
                <tr>
                    <StyledCol><h3 key = {event.id}>{event.title}</h3></StyledCol>
                    <StyledCol><h3>{event.startdate}</h3></StyledCol>
                    <StyledCol><Button variant='primary' href= {'/event'}>Go</Button></StyledCol>
                </tr>)}
            </div>
        );
    }
}