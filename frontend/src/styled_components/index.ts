import styled from "styled-components"

export const Flex = styled.div`
    display: flex;
    gap: 10px 10px;
`

export const FlexColumn = styled.div`
    display: flex;
    flex-direction: column;
    gap: 10px 10px;
`

export type FloatingMenuProps = {
  show?: boolean
}

export const FloatingMenu = styled.div<FloatingMenuProps>`
    position: absolute;
    top: 0;
    left: -62px;

    @media (min-width: 1400px) {
      left: -67px;
    }

    z-index: 99;
    transition: 0.3s ease all;
    ${props => props.show ? 'top: 20px;' : 'top: -300px;'}
    
    > * {
        background-color: #fff;
        position: fixed;
        border-radius: 6px;
    }
`

export const AppContainer = styled.main`
  width: 100%;
  height: 100vh;
  background-color: #eee;
  display: flex;
  justify-content: center;
  overflow: auto;
`

export const BackButtonContainer = styled.div`
  position: fixed;
  display: flex;
  bottom: 20px;
  left: 20px;

  > button {
    box-shadow: 0 0 10px 7px #edebe9;

    svg {
      transform: scale(0.7);
      transition: 0.2s;
    }

    &:hover {
      svg {
        transform: scale(0.5);
      }
    }
  }
`

export const AppContainer2 = styled.div`
  max-width: 960px;
  @media (min-width: 1400px) {
    max-width: 1125px;
  }

  width: 100%;
  padding: 15px;
  font-size: 14px;
  background-color: #ffff;
  display: flex;
  flex-direction: column;
  height: max-content;
  min-height: 100vh;
  box-shadow: 0 0 50px 3px #edebe9;
  position: relative;
  border: 1px solid #e1dfdd;
`
