import './App.css';
import {useFiles} from "./hooks/files";
import {Loader} from "./components/Loader";
import {ErrorMessage} from "./components/ErrorMessage";
import {File} from "./components/File";
import {v4 as uuidv4} from 'uuid';

// todo: use React Admin: https://marmelab.com/react-admin
function App() {
    const {loading, error, files} = useFiles()

    return (
        <div>
            {loading && <Loader/>}
            {error && <ErrorMessage error={error}/>}
            {files.map(file => <File file={file} key={uuidv4()}/>)}
        </div>
    );
}

export default App;
