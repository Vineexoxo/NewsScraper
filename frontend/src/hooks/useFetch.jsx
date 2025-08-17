import { Request } from "@/request/request";
import {useState, useEffect, useCallback} from "react";
// import axios from "axios";
export function useFetch(url, params={}, auto=false){
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const fetchData=useCallback(async ()=>{
        try{
            setLoading(true);
            setError(null);
            const response = await Request.post(url, params);
            console.log("response", response);
            setData(response.data);
        }catch(err){
            setError(err);
        }finally{
            setLoading(false);
        }
    }, [url, JSON.stringify(params)]);
    
    useEffect(()=>{
        if(url==="") return;
        if(auto) fetchData();
        
    },[url, JSON.stringify(params)]);
    

    return {data, loading, error, fetchData};
}
