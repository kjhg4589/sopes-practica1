import React from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Card from '@mui/material/Card';

const Historial= (props) => {
    return (
    <Card 
        sx={{ maxWidth: 800, marginTop: 5}}
        >
        <TableContainer 
            sx={{maxHeight: 600}}
            component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
                <TableHead>
                <TableRow>
                    <TableCell align="right">Numero 1</TableCell>
                    <TableCell align="right">Numero 2</TableCell>
                    <TableCell align="center">Operacion</TableCell>
                    <TableCell align="right">Resultado</TableCell>
                    <TableCell align="center">Fecha</TableCell>
                </TableRow>
                </TableHead>
                <TableBody>
                {props.historial?.map((row) => (
                    <TableRow
                        key={row.id}
                        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                    >
                    <TableCell align="right">{row.numero1}</TableCell>
                    <TableCell align="right">{row.numero2}</TableCell>
                    <TableCell align="right">{row.operacion}</TableCell>
                    <TableCell align="right">{row.resultado}</TableCell>
                    <TableCell align="right">{row.Fecha}</TableCell>
                    </TableRow>
                ))}
                </TableBody>
            </Table>
        </TableContainer>
    </Card>
    );
}

export default Historial;