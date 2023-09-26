import {Datagrid, List, TextField} from "react-admin";

export interface File {
    id: number;
    filepath: string;
    created_at: Date;
}

export const FileList = () => (
    <List>
        <Datagrid rowClick="show">
            <TextField source="id"/>
            <TextField source="filepath"/>
            <TextField source="created_at"/>
        </Datagrid>
    </List>
);