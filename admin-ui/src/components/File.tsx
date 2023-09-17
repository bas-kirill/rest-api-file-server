interface FileProps {
    file: string
}

export function File({file}: FileProps) {
    console.log(file)
    return (
        <>
            <div>
                Filename: {file}
            </div>
        </>
    )
}