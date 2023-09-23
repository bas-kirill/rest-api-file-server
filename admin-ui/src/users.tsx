import {Datagrid, List, SimpleList, TextField} from "react-admin";
import {Theme, useMediaQuery} from "@mui/material";

export const UserList = () => {
    const isSmall = useMediaQuery<Theme>((theme) => theme.breakpoints.down("sm"));
    return (
        <List>
            {isSmall ? (
                <SimpleList
                    primaryText={(record) => record.name}
                    secondaryText={(record) => record.username}
                    tertiaryText={(record) => record.email}
                />
            ) : (
                <Datagrid rowClick="edit">
                    <TextField source="id"/>
                    <TextField source="name"/>
                    <TextField source="username"/>
                    <TextField source="email"/>
                    <TextField source="address.street"/>
                    <TextField source="phone"/>
                    <TextField source="website"/>
                    <TextField source="comapny.name"/>
                </Datagrid>
            )}
        </List>
    )
};