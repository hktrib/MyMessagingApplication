import Loading from "../Loading";
import { useState } from "react";
import { useLocation } from "react-router-dom";

const VerifyEmail = () => {
    
    const [isLoading, setIsLoading] = useState(true)
    const location = useLocation()

    const queryParams = new URLSearchParams(location.search)

    console.log(queryParams.get("email"))
    console.log(queryParams.get("secret_code"))

    return (
        <div className="VerifyEmail">
            <p>Give us a second while we verify your credentials</p>   
        </div>
    )
}


export default VerifyEmail;
