import React, { Component, Fragment } from "react";
import axios from 'axios';
import Historial from "./Historial";
import { Card, TextField, Select, MenuItem, Button} from "@mui/material";
import api_host from "./Creds";

class Calculadora extends Component {

    state = {
        hiss: [],
        num1: Number,
        num2: Number,
        operacion: '+',
        resultado: Number,
    }

    componentDidMount() {
        this.updateHiss();
    }

    updateHiss() {
        axios.get(`${api_host}/oper/report`)
        .then(result=>{
            const hiss = result.data
            this.setState({
                hiss
            })
        }).catch(console.log)
    }

    handleChangeNum1(value) {
        this.setState ({
            num1: value
        })
    }

    handleChangeNum2(value) {
        this.setState ({
            num2: value
        })
    }


    handleChangeOperacion(value) {
        this.setState ({
            operacion: value
        })
    }

    submitForm = () => {
        
        const {num1, num2, operacion} = this.state;
        axios.post(`${api_host}/oper`, {numero1: num1, numero2: num2, operacion: operacion})
            .then(result => {
                this.setState({
                    resultado: result.data.resultado
                })
            }).catch(console.log)
        this.updateHiss()
    }

    render() {
        
        const datos  = this.state.hiss
    
        return (
            <Fragment>
                <div>
                    <Card sx={{marginTop: 15}}>
                        <TextField sx={{width: 100, margin: 2}} 
                                    value={this.state.num1}
                                    onChange={(env)=>{this.handleChangeNum1(env.target.value)}}
                                    id="num1" 
                                    label="Numero 1" 
                                    type="text"
                                    variant="outlined" />
                        <Select
                            sx={{width: 75, margin: 2}}
                            labelId="demo-simple-select-label"
                            id="demo-simple-select"
                            value={this.state.operacion}
                            onChange={(env)=>{this.handleChangeOperacion(env.target.value)}}
                            label="operacion"
                        >
                            <MenuItem value={"+"}>+</MenuItem>
                            <MenuItem value={"-"}>-</MenuItem>
                            <MenuItem value={"*"}>*</MenuItem>
                            <MenuItem value={"/"}>/</MenuItem>
                        </Select>
                        <TextField sx={{width: 100, margin: 2}} 
                                    value={this.state.num2}
                                    onChange={(env)=>{this.handleChangeNum2(env.target.value)}}
                                    id="num2" 
                                    label="Numero 2" 
                                    variant="outlined" />
                        <Button sx={{width: 75, margin: 2}} variant="contained" onClick={this.submitForm}>=</Button>

                        <TextField sx={{width: 100, margin: 2}} 
                                    value={this.state.resultado}                                                          
                                    id="result" 
                                    label="Resultado" 
                                    variant="outlined" />
                        
                    </Card>
                    <Historial historial={datos} />
                </div>
            </Fragment>              
        )
    }
}

export default Calculadora;  