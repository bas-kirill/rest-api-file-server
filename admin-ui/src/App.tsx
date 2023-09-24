import {Admin, ListGuesser, Resource, ShowGuesser,} from "react-admin";
import {dataProvider} from "./dataProvider";
import FileOpenIcon from '@mui/icons-material/FileOpen';
import {Dashboard} from "./Dashboard";
import {FileList} from "./Files";
import {DebugLayout} from "./DebugLayout";

export const App = () => (
    <Admin dataProvider={dataProvider} dashboard={Dashboard}>
        <Resource
            name="file"
            show={ShowGuesser}
            list={FileList}
            icon={FileOpenIcon}/>
    </Admin>
);
