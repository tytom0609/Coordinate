import {Button} from 'react-bootstrap'
import React, { Component } from 'react'
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

export class EventPage extends Component {
    state = {
        requests: []
    }

    constructor() {
        super();
        api.get('/Requests').then(res => {
            console.log(res.data)
            this.setState({requests : res.data})
        })
    }

    signUp() {
        alert('Successfully signed up!')
    }

    render(){
        return(
            <div>
                <div class='row'><img src="https://placekitten.com/150/150" roundedCircle></img>
                <div class = 'col'><h1>Cat Adoption Drive:</h1><p>We have lots of lovely kitties that are looking for their forever homes! Come to this small party
                     where you will get a chance to fall in love with all of our felines! We need assistance in making sure the cats aren't overwhelmed and small party
                      favors to entertain potential adoptees during this event. Our building is located at 3400 Northwood Lane,
                    Indianapolis, IN, 46268!</p></div></div>
                <div class='row'><h3>We Need:</h3></div>
                <tr>
                    <StyledHead>Request</StyledHead>
                    <StyledHead>Info</StyledHead>
                    <StyledHead>Status</StyledHead>
                    <StyledHead>Sign Up</StyledHead>
                </tr>
                {this.state.requests.map(request => 
                <tr>
                    <StyledCol><h3 key = {request.id}>{request.title}</h3></StyledCol>
                    <StyledCol><h3>{request.description}</h3></StyledCol>
                    <StyledCol><h3>{request.quantityobtained}/{request.quantityneeded} Signed Up</h3></StyledCol>
                    <StyledCol><Button onClick = {this.signUp} variant='primary'>Sign Up</Button></StyledCol>
                    
                </tr>)}
            </div>
        );
    }
}