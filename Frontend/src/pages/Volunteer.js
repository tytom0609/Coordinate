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

export class Volunteer extends Component {
    state = {
        nonprofits: []
    }

    constructor() {
        super();
        api.get('/NonProfits').then(res => {
            console.log(res.data)
            this.setState({nonprofits : res.data})
        })
    }

    render(){
        return(
            <div>
                <div class='row'><h3>Non-Profits in your Area:</h3></div>
                <tr>
                    <StyledHead>Non-Profit</StyledHead>
                    <StyledHead>Go to Page</StyledHead>
                </tr>
                {this.state.nonprofits.map(nonprofit => 
                <tr>
                    <StyledCol><h3 key = {nonprofit.id}>{nonprofit.name}</h3></StyledCol>
                    <StyledCol><Button variant='primary' href= {'/nonprofit'}>Go</Button></StyledCol>
                </tr>)}
            </div>
        );
    }
}
