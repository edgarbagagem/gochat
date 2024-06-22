export const isAuthenticated = () => {
  return !!sessionStorage.getItem("jwtToken");
};
