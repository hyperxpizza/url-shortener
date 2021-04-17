import React, {useState, useEffect} from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';

import Container from '@material-ui/core/Container';
import Typography from '@material-ui/core/Typography';

const Info = () => {
    const {id} = useParams();
    
    const [data, setData] = useState({});

    useEffect(() => {

        let endpoint = process.env.REACT_APP_SERVER_HOST + ":" + "api" + process.env.REACT_APP_SERVER_PORT + "/" + id + process.env.REACT_APP_INFO_ENDPOINT;
        axios.get(endpoint)
            .then((response) => {
                setData(response.data);
            }, (error) => {
                console.log(error);
            });
            
    }, [setData]);
    
    return(
        <div>
            <Container>
                <Typography variant="h4" component="h4" align="center">
                    Redirect to: <a href={data.url}>{data.url}</a>
                </Typography>

                <Typography variant="h4" component="h4" align="center">
                    Hits: {data.hits}
                </Typography>
            </Container>
        </div>
    );
} 



export default Info;