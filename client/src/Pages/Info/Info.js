import React, {useState, useEffect} from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';

import Container from '@material-ui/core/Container';

const Info = () => {
    const {id} = useParams();
    
    const [data, setData] = useState();

    useEffect(() => {

        let endpoint = process.env.REACT_APP_SERVER_HOST + ":" + process.env.REACT_APP_SERVER_PORT + "/" + id + process.env.REACT_APP_INFO_ENDPOINT;


        axios.get(endpoint)
            .then((response) => {
                setData(response.data);
            }, (error) => {
                console.log(error);
            });
            
    }, [setData]);
    
    return(
        <div>
            <h1>Info</h1>
        </div>
    );
} 



export default Info;