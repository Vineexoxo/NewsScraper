"use client";

import * as React from "react";

import { Progress } from "@/components/retroui/Progress";
import { Text } from "@/components/retroui/Text";
import { Button } from "../retroui/Button";
import { Dialog } from "@radix-ui/react-dialog";
import { Link, Navigate } from "react-router-dom";
export default function ProgressBar({loading, data}) {
    const [progress, setProgress] = React.useState(13);

    React.useEffect(() => {
        if(loading){
            const timer = setTimeout(() => setProgress(progress+10), 500);
            return () => clearTimeout(timer);
        }
    }, [progress]);
    React.useEffect(()=>{
        if(!loading){
            setProgress(100);
                    
        }
    },[loading])
    return (
        <div>
            <Progress value={progress} className="w-[60%] mt-4 mb-6 mx-auto" />
            <Dialog.Footer>
                <Dialog.Trigger asChild className="m-4">

                    {loading ?
                        <Text>
                            Loading...
                        </Text>
                        :
                        <Button asChild onClick={() => {
                            
                        }}>
                            <Link 
                            to={`/articles/${encodeURIComponent(data.link)}`}>
                                Done
                            </Link>
                        </Button>
                    }
                </Dialog.Trigger>
            </Dialog.Footer>
        </div>
    )

    
}