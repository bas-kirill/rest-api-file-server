import {Layout} from 'react-admin';
import {ReactQueryDevtools} from 'react-query/devtools';

export const DebugLayout = props => (
    <>
        <Layout {...props} />
        <ReactQueryDevtools initialIsOpen={false}/>
    </>
);