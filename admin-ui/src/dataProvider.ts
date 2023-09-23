import fakeRestDataProvider from "ra-data-fakerest";
import data from "./data.json";
import jsonServerProvider from "ra-data-json-server";

export const dataProvider = jsonServerProvider(
    import.meta.env.VITE_JSON_SERVER_URL
)