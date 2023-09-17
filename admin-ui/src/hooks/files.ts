import {useEffect, useState} from "react";
import axios, {AxiosError} from "axios";

export function useFiles() {
    const [files, setFiles] = useState<string[]>([])
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")

    async function fetchListFiles() {
        try {
            setError("")
            setLoading(true)
            const response = await axios.get<string[]>("http://localhost:8080")
            console.log(response.data)
            setFiles(response.data)
            setLoading(false)
        } catch (e: unknown) {
            const error = e as AxiosError
            setLoading(false)
            setError(error.message)
        }
    }

    useEffect(() => {
        fetchListFiles()
    }, []);

    return {loading, error, files}
}