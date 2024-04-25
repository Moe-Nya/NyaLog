//时间处理
function formatDate(dateTime) {
    const date = new Date(dateTime);
    const formattedDate = date.toISOString().split('T')[0];
    return formattedDate;
}

export default formatDate